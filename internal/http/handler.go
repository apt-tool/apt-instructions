package http

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/ptaas-tool/ftp-server/internal/storage"

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

func (h Handler) Upload(ctx *fiber.Ctx) error {
	name := ctx.Query("name")

	file, err := ctx.FormFile("script")
	if err != nil {
		log.Println(fmt.Errorf("[handler.Upload] failed to get file error=%w", err))

		return fiber.ErrBadRequest
	}

	return ctx.SaveFile(file, fmt.Sprintf("./data/attacks/%s.sh", name))
}

func (h Handler) List(ctx *fiber.Ctx) error {
	entries, err := os.ReadDir("./data/attacks/")
	if err != nil {
		log.Println(fmt.Errorf("[handler.List] failed to get files error=%w", err))

		return fiber.ErrInternalServerError
	}

	list := make([]string, 0)

	for _, e := range entries {
		if e.Name() != "NOTE.txt" {
			list = append(list, strings.Replace(e.Name(), ".sh", "", 1))
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

	path := fmt.Sprintf("./data/attacks/%s.sh", req.Path)

	code := 0
	cmd, err := exec.Command("/bin/sh", path, req.Param).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			code = exitError.ExitCode()
		} else {
			log.Println(fmt.Errorf("[handler.Execute] failed to get files error=%w", err))

			return fiber.ErrInternalServerError
		}
	}

	path = fmt.Sprintf("./data/docs/%d.txt", req.DocumentID)
	f, err := os.Create(path)
	if err != nil {
		log.Println(fmt.Errorf("[handler.Execute] failed to store log file error=%w", err))

		return fiber.ErrInternalServerError
	}

	defer func(f *os.File) {
		er := f.Close()
		if er != nil {
			log.Println(fmt.Errorf("[handler.Execute] failed to close file error=%w", er))
		}
	}(f)

	_, _ = f.Write(cmd)

	if er := h.MinioClient.Put(fmt.Sprintf("%d.txt", req.DocumentID), path); er != nil {
		log.Println("[handler.Execute] failed to store file error=%w", err)

		return fiber.ErrInternalServerError
	}

	if er := os.Remove(path); er != nil {
		log.Println("[handler.Execute] failed to remove local file error=%w", er)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": code,
	})
}
