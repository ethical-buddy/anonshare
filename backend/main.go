package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/muskiteer/anonshare/backend/database"
	"github.com/muskiteer/anonshare/backend/routes"
)

func main() {
	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, proceeding with system environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Println("No PORT environment variable detected, defaulting to 5000")
	}

	// Initialize database
	db := database.InitDatabase()
	log.Printf("✅ Initialized database successfully: %s", db.Name())

	// Setup HTTP routes
	mux := http.NewServeMux()
	handler := routes.SetupRoutes(mux, db)

	addr := "127.0.0.1:" + port
	log.Printf("🚀 Server is running on http://%s", addr)

	// Start server
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
