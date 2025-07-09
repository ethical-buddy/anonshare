package models

import (
	"gorm.io/gorm"
)

type PeerInfo struct {
	gorm.Model
	Peer_ID     string `json:"peer_id"`
	IP          string `json:"ip"`
	Port        string `json:"port"`
	FilePath    string `json:"file_path"`
	Filename    string `json:"filename"`
	Description string `json:"description"`

	// Foreign key to FileMetadata
	FileMetadataID uint         `json:"file_metadata_id" gorm:"not null"`
	FileMetadata   FileMetadata `json:"file_metadata" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type PeerMetadata struct { // PeerMetadata represents the metadata of a peer in the system
	gorm.Model
	Peer_ID     string `json:"peer_id"`
	IP          string `json:"ip"`
	Port        string `json:"port"`
	FilePath    string `json:"file_path"`
	Filename    string `json:"filename"`
	Description string `json:"description"`
}
