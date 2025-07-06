package upload

import (
	"log"
	"os"
	"time"
	"errors"

	"github.com/muskiteer/anonshare/internal"
	"github.com/muskiteer/anonshare/models"
)

func HandleUpload(filePath string, description string,Peer_Port string) (models.FileMetadata,error) {

	info,err := os.Stat(filePath)
	if err != nil {
		log.Printf("Error accessing file: %v on upload.go", err)
		return models.FileMetadata{}, err
	}

	if info.IsDir() {
		log.Println("Error: Uploading directories is not supported currently")
		return models.FileMetadata{}, errors.New("uploading directories is not supported")
	}

	if info.Size() == 0 {
		log.Println("Error: Cannot upload an empty file")
		return models.FileMetadata{}, errors.New("cannot upload an empty file")
	}

	hash_file, err := internal.CalculateFileHash(filePath)
	if err != nil {
		log.Printf("Error calculating file hash: %v", err)
		return models.FileMetadata{}, errors.New("error calculating file hash")
	}
	file_type, err := internal.DetectFileType(filePath)
	if err != nil {
		log.Printf("Error detecting file type: %v", err)
		return models.FileMetadata{}, errors.New("error detecting file type")
	}

	PeerInfo := models.PeerInfo{
		Peer_ID: internal.GetOrCreateNodeID(),
		Filename:   info.Name(),
		Description: description,
		IP:        internal.GetLocalIP(),
		Port:      Peer_Port,
		FilePath:  filePath,

	};
		
	
	Filemetadata:= models.FileMetadata{
		Hash:        hash_file,
		Size:        internal.FormatFileSize(info.Size()),
		Time:         info.ModTime().Format(time.RFC3339),
		Type:        file_type,
		Peers:    []models.PeerInfo{PeerInfo},
	};



	return Filemetadata,nil
}
