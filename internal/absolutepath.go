package internal

import (
	"os"
	"path/filepath"
)

func GetAbsolutePath(filePath string) (string, error) {
    if filepath.IsAbs(filePath) {
        return filePath, nil
    }
    
    // Get executable directory
    ex, err := os.Executable()
    if err != nil {
        return "", err
    }
    exPath := filepath.Dir(ex)
    
    // Resolve relative to executable
    return filepath.Join(exPath, filePath), nil
}