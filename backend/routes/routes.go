package routes

import (
	"net/http"
	"gorm.io/gorm"
	"github.com/muskiteer/anonshare/backend/controllers"
)

func SetupRoutes(r *http.ServeMux, db *gorm.DB) http.Handler {
	r.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Received request for /upload"))
		controllers.UploadHandler(w, r, db)
	})
	r.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		controllers.GettingFilesHandler(w, r, db)
	})

	r.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		controllers.DownloadHandler(w, r, db)
	})
	return r
}