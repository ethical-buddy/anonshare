package models

// import "path/filepath"

type FileMetadata struct {
	Hash  string   `json:"hash" gorm:"uniqueIndex"`
	Size  string    `json:"size"`
	Time  string   `json:"time"`
	Type string   `json:"type"`

	Peers []PeerInfo `json:"peers"`
}





