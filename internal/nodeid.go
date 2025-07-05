package internal

import (
	"os"
	"path/filepath"
	"github.com/google/uuid"
	"log"
)

// GetOrCreateNodeID ensures a persistent .nodeid file exists and returns the NodeID
func GetOrCreateNodeID() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("Failed to get executable path:", err)
	}
	nodeIDPath := filepath.Join(filepath.Dir(exePath), ".nodeid")

	// Check if file exists and is not empty
	if data, err := os.ReadFile(nodeIDPath); err == nil && len(data) > 0 {
		return string(data)
	}

	// Generate new UUID and save it
	id := uuid.New().String()
	if err := os.WriteFile(nodeIDPath, []byte(id), 0644); err != nil {
		log.Fatal("Failed to write .nodeid file:", err)
	}
	return id
}
