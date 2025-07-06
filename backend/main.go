package main

import (
	"log"
	// "net/http"
	"os"

	// "github.com/muskiteer/anonshare/internal"
	"github.com/joho/godotenv"
	"github.com/muskiteer/anonshare/backend/database"
)

func main() {
	err := godotenv.Load() // Load the .env file from the current directory
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	Port := os.Getenv("PORT")

	if Port == "" {
		Port = "5000"
		log.Println("No PORT environment variable detected, defaulting to 5000")
	}
	database.InitDatabase()
	log.Printf("initialized database successfully, running backend server on port %s", Port)

	
}