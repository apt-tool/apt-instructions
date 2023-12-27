package http

import (
	"fmt"
	"github.com/ptaas-tool/ftp-server/internal/storage"
	"log"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	AccessKey   string
	PrivateKey  string
	MinioClient storage.Client
}

func (h Handler) Health(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) Download(ctx *fiber.Ctx) error {
	path := ctx.Query("path")

	url, err := h.MinioClient.Get(fmt.Sprintf("%s.txt", path))
	if err != nil {
		log.Println(fmt.Errorf("[handler.Download] failed to get url error=%w", err))

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).SendString(url)
}

func (h Handler) List(ctx *fiber.Ctx) error {
	entries, err := os.ReadDir("./libatks/")
	if err != nil {
		log.Println(fmt.Errorf("[handler.List] failed to get files error=%w", err))

		return fiber.ErrInternalServerError
	}

	list := make([]string, 0)

	for _, e := range entries {
		if e.Name() != "go.mod" && e.Name() != "go.sum" {
			list = append(list, e.Name())
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(list)
}

func (h Handler) Execute(ctx *fiber.Ctx) error {
	req := new(ExecuteRequest)

	if err := ctx.BodyParser(&req); err != nil {
		log.Println(fmt.Errorf("[handler.Execute] failed to parse body error=%w", err))

		return fiber.ErrBadRequest
	}

	path := fmt.Sprintf("libatks/%s/main.go", req.Path)
	code := 0

	// command to execute your Golang script
	cmd := exec.Command("go", "run", path)

	// redirect output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(fmt.Errorf("[handler.Execute] error creating StdoutPipe error=%w", err))

		return fiber.ErrInternalServerError
	}

	// start the command
	if err := cmd.Start(); err != nil {
		log.Println(fmt.Errorf("[handler.Execute] Error starting command error=%w", err))

		return fiber.ErrInternalServerError
	}

	// read the output
	output := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := stdout.Read(buf)
		if err != nil {
			break
		}

		output = append(output, buf[:n]...)
	}

	// wait for command to finish
	if err := cmd.Wait(); err != nil {
		exitError, ok := err.(*exec.ExitError)
		if !ok {
			log.Println(fmt.Errorf("[handler.Execute] command execution error=%w", err))

			return fiber.ErrInternalServerError
		}

		// get the exit code
		code = exitError.ExitCode()
	}

	newPath := fmt.Sprintf("./data/docs/%d.txt", req.DocumentID)

	f, err := os.Create(newPath)
	if err != nil {
		log.Println(fmt.Errorf("[handler.Execute] failed to store log file error=%w", err))

		return fiber.ErrInternalServerError
	}

	_, _ = f.Write(output)

	if er := h.MinioClient.Put(fmt.Sprintf("%d.txt", req.DocumentID), newPath); er != nil {
		log.Println(fmt.Errorf("[handler.Execute] failed to store file error=%w", err))

		return fiber.ErrInternalServerError
	}

	if er := os.Remove(newPath); er != nil {
		log.Println(fmt.Errorf("[handler.Execute] failed to remove local file error=%w", er))
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": code,
	})
}
