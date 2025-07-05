package upload

import (
	"log"
	"os"
	"time"

	"github.com/muskiteer/anonshare/internal"
	"github.com/muskiteer/anonshare/models"
)

func HandleUpload(filePath string, description string,Peer_Port string) error {

	info,err := os.Stat(filePath)
	if err != nil {
		log.Printf("Error accessing file: %v on upload.go", err)
		return err
	}

	if info.IsDir() {
		log.Println("Error: Uploading directories is not supported currently")
		return nil
	}

	if info.Size() == 0 {
		log.Println("Error: Cannot upload an empty file")
		return nil
	}

	hash_file, err := internal.CalculateFileHash(filePath)
	if err != nil {
		log.Printf("Error calculating file hash: %v", err)
		return err
	}
	file_type, err := internal.DetectFileType(filePath)
	if err != nil {
		log.Printf("Error detecting file type: %v", err)
		return err
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

	
	return nil
}
