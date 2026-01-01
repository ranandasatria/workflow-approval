package repository

import (
	"fmt"
	"log"
	"os"
	"workflow-approval/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("database connected")

	runMigrations(db)

	return db
}

func runMigrations(db *gorm.DB) {
	log.Println("running database migrations...")
	err := db.AutoMigrate(
		&model.Workflow{},
		&model.WorkflowStep{},
		&model.Request{},
	)
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}