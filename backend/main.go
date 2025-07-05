package main

import (
	"log"
	"net/http"
	"os"

	"github.com/muskiteer/anonshare/internal"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	Port := os.Getenv("PORT")

	if Port == "" {
		Port = "5000"
		log.Println("No PORT environment variable detected, defaulting to 5000")
	}

	// log.Fatal(http.ListenAndServe(":"+Port, nil))
}