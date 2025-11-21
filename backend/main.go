package main

import (
	"log"

	"backend/database"
	"backend/routes"
)

func main() {
	// Initialize DB and auto-migrate models
	if err := database.Init(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer database.Close()

	// Setup routes and run server
	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
