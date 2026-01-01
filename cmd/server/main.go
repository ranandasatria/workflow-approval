package main

import (
	"log"
	"workflow-approval/internal/app"
	"workflow-approval/internal/repository"
	"workflow-approval/internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db := repository.NewPostgresDB()

	container := app.NewContainer(db)

	fiberApp := fiber.New()

	router.SetupRoutes(fiberApp, container.Handlers)

	log.Fatal(fiberApp.Listen(":9000"))
}