package main

import (
	"log"

	"github.com/minamoto-m/project-manager-backend/internal/domains"
	"github.com/minamoto-m/project-manager-backend/internal/infrastructures"
)

func main() {
	db := infrastructures.SetupDB()

	if err := db.AutoMigrate(&domains.Task{}); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
	log.Println("Completed migrations")
}
