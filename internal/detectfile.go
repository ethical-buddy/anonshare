package internal

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DetectFileType(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	buffer := make([]byte, 512)
	n, err := f.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer[:n])

	// Fallback using extension
	if contentType == "application/octet-stream" {
		ext := filepath.Ext(filePath)
		mimeType := mime.TypeByExtension(ext)
		if mimeType != "" {
			contentType = mimeType
		}
	}

	parts := strings.Split(contentType, "/")
	if len(parts) == 0 {
		return "application", nil
	}
	mainType := parts[0]
	ext := strings.ToLower(filepath.Ext(filePath))

	// 📘 Classify text and document-like types
	if mainType == "text" ||
		strings.Contains(contentType, "pdf") ||
		strings.Contains(contentType, "epub") ||
		strings.Contains(contentType, "msword") ||
		strings.Contains(contentType, "officedocument") ||
		strings.Contains(contentType, "rtf") ||
		strings.Contains(contentType, "vnd.oasis.opendocument") ||
		ext == ".pdf" || ext == ".epub" || ext == ".docx" || ext == ".odt" || ext == ".txt" {
		return "document", nil
	}

	// Acceptable main types
	switch mainType {
	case "image", "audio", "video":
		return mainType, nil
	}

	return "application", nil
}


