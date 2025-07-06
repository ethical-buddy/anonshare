package routes

import (
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/files", gettingFilesHandler)
}