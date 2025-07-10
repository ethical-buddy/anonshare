package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"gorm.io/gorm"
	"github.com/muskiteer/anonshare/utils"

	"github.com/muskiteer/anonshare/models"
)

func UploadHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	if r.Method != http.MethodPost {
		utils.JSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		log.Println("Method not allowed in UploadHandler")
		return
	}

	var fileMetadata models.FileMetadata
	if err := json.NewDecoder(r.Body).Decode(&fileMetadata); err != nil {
		log.Printf("Error parsing request body: %v in UploadHandler", err)
		utils.JSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if fileMetadata.Hash == "" || fileMetadata.Size == "" || len(fileMetadata.Peers) == 0 {
		log.Println("Error: Missing required fields in file metadata")
		utils.JSONError(w, http.StatusBadRequest, "Missing required fields")
		return
	}
	
	if err := models.UploadingInDB(db, &fileMetadata); err != nil {
		
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, "File metadata and peer created successfully")
	
	
}

func GettingFilesHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var fileMetadata []models.FileMetadata
	if err := models.GettingFilesFromDB(db, &fileMetadata); err != nil {
		log.Printf("Error retrieving files from database: %v in GettingFilesHandler", err)
		utils.JSONError(w, http.StatusInternalServerError, "Failed to retrieve files")
		return
	}
	
	utils.JSONResponse(w, http.StatusOK, fileMetadata)
	
}

func DownloadHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type QueryParams struct {
		Hash string `json:"hash"`
	}
	var params QueryParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Printf("Error parsing query parameters: %v in DownloadHandler", err)
		utils.JSONError(w, http.StatusBadRequest, "Invalid query parameters")
		return
	}
	if params.Hash == "" {
		utils.JSONError(w, http.StatusBadRequest, "Missing file hash")
		log.Println("Error: Missing file hash in DownloadHandler")
		return
	}

	

	peermetadata, err := models.GettingPeersFromDB(db, params.Hash)

	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		log.Printf("Error retrieving peers for hash %s: %v in DownloadHandler", params.Hash, err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, peermetadata)
}
