package cli

import (
	"flag"
	"log"
	"path/filepath"

	// "github.com/muskiteer/anonshare/internal"
	
	"github.com/muskiteer/anonshare/peer/upload"
)

func Cli(peerPort string) {
	// Define CLI flags
	uploadPath := flag.String("u", "", "File path to upload")
	description := flag.String("d", "", "File description")
	flag.Parse()

	if *uploadPath == "" || *description == "" {
		log.Fatal("Usage: ./anonshare -u <file_path> -d <description>")
	}

	// Get absolute file path
	absolutePath, err := filepath.Abs(*uploadPath)
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}
	
	err = upload.HandleUpload(absolutePath, *description, peerPort)
	if err != nil {
		log.Fatalf("Upload failed: %v", err)
	}
}
