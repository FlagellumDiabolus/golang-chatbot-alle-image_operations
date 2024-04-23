package main

import (
	"golang-chatbot-alle-image_operations/internal/database"
	"golang-chatbot-alle-image_operations/internal/server"
	"log"
	"net/http"
)

func main() {
	err := database.InitializeDB("data/database.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	// Initialize HTTP server
	srv := server.NewServer()

	// Set up routes
	srv.SetupRoutes()

	// Start HTTP server
	log.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", srv)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
