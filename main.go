package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ptaas-tool/ftp-server/internal/http"
	"github.com/ptaas-tool/ftp-server/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// get env variables
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	private := os.Getenv("PRIVATE_KEY")
	access := os.Getenv("ACCESS_KEY")

	// get minio configs
	minioCli, err := storage.New(storage.LoadConfig(os.Getenv("MINIO_CLUSTER")))
	if err != nil {
		panic(err)
	}

	// create new fiber app
	app := fiber.New()

	app.Use(cors.New())

	// create new handler
	h := http.Handler{
		AccessKey:   access,
		PrivateKey:  private,
		MinioClient: minioCli,
	}

	app.Get("/health", h.Health)
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Get("/", h.List)
	app.Get("/download", h.AccessMiddleware, h.Download)
	app.Post("/upload", h.Upload)
	app.Post("/execute", h.AuthMiddleware, h.Execute)

	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(fmt.Errorf("failed to start ftp server error=%w", err))
	}
}
