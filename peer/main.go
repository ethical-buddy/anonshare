package main 

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "net/http"
    "github.com/muskiteer/anonshare/peer/cmd"
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
    
    // Start HTTP server in background for peer-to-peer file transfers
    go func() {
        log.Printf("🌐 Starting peer server on port %s for file transfers", PeerPort)
        if err := http.ListenAndServe("0.0.0.0:"+PeerPort, nil); err != nil {
            log.Printf("❌ Error starting peer server: %v", err)
        }
    }()
    
    log.Printf("🚀 Peer client ready! Port: %s, Backend: %s", PeerPort, BackendUrl)
    
    // Start the interactive CLI
    cli.Cli()
}