package main

import (
	"log"

	"workflow-approval/internal/repository"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db := repository.NewPostgresDB()
	_ = db

	log.Println("server started")
}
