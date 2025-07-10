package functions

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/muskiteer/anonshare/models"
	"github.com/muskiteer/anonshare/peer/upload"
)

func uploading_the_file(filepath string,description string,Port string) {
	BackendUrl := os.Getenv("BACKEND_URL")
	if BackendUrl == "" {
		log.Println("no BACKEND_URL environment variable detected")
		return
	}
	url := BackendUrl + "/upload"

	var filemetadata models.FileMetadata
	filemetadata, err := upload.HandleUpload(filepath, description, Port)
	if err != nil {
		log.Println("error handling upload: " + err.Error())
		return
	}
	jsonPayload, err := json.Marshal(filemetadata)
	if err != nil {
		log.Println("error marshalling file metadata: " + err.Error())
		return
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("error creating request: " + err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error sending request to backend: " + err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
        log.Printf("error uploading file to backend: %s", resp.Status)
		return
    }
	log.Println("File uploaded successfully to backend")
}

func gettings_files_info() {
	BackendUrl := os.Getenv("BACKEND_URL")
	if BackendUrl == "" {
		log.Println("no BACKEND_URL environment variable detected")
		return
	}
	url := BackendUrl + "/files"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("error fetching files from backend: " + err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("error fetching files from backend: " + resp.Status)
	}

	var filemetadata []models.FileMetadata
	if err := json.NewDecoder(resp.Body).Decode(&filemetadata); err != nil {
		log.Println("error decoding response from backend: " + err.Error())
		return
	}

	if len(filemetadata) == 0 {
		log.Println("No files found in backend")
		return
	}

	log.Printf("Successfully fetched %d files from backend", len(filemetadata))

	for _, file := range filemetadata {
		log.Printf("File: %s, Size: %s, Type: %s, Hash: %s, Peers: %d",
			"will add later to see the name from first user", file.Size, file.Type, file.Hash, len(file.Peers))
	}

}

func file_download(hash string) {
	BackendUrl := os.Getenv("BACKEND_URL")
	if BackendUrl == "" {
		log.Println("no BACKEND_URL environment variable detected")
		return
	}
	if hash == "" {
		log.Println("no hash provided for getting peers")
		return
	}
	url := BackendUrl + "/download"

	payload := map[string]string{"hash": hash}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println("error marshalling payload: " + err.Error())
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("error creating request: " + err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error sending request to backend: " + err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("error fetching peers from backend: " + resp.Status)
		return
	}
	var peers []models.PeerInfo
	if err := json.NewDecoder(resp.Body).Decode(&peers); err != nil {
		log.Println("error decoding response from backend: " + err.Error())
		return
	}
	if len(peers) == 0 {
		log.Println("No peers found for the given hash")
		return
	}

	downloading_from_Peers(peers)
	
}

func downloading_from_Peers(peers []models.PeerInfo) {
	log.Printf("Downloading the file will be added soon")
	for _, peer := range peers {
		log.Printf("Peer ID: %s, IP: %s, Port: %s, File Path: %s, Filename: %s, Description: %s",
			peer.Peer_ID, peer.IP, peer.Port, peer.FilePath, peer.Filename, peer.Description)
		// Here you would implement the logic to download the file from the peer
	}
}
