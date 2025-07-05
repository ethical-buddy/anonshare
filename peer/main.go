package main 

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PeerPort := os.Getenv("PEER_PORT")
	if PeerPort == "" {
		PeerPort = "8080"
		log.Println("No PEER_PORT environment variable detected, defaulting to 8080")
	}
	BackendUrl := os.Getenv("BACKEND_URL")
	if BackendUrl == "" {
		log.Fatal("No BACKEND_URL environment variable detected")
	}
	log.Printf("Peer server running on port %s, connecting to backend at %s", PeerPort, BackendUrl)


}


