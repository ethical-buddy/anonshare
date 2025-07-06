package models

import "gorm.io/gorm"

type FileMetadata struct {
	gorm.Model
	Hash string `json:"hash" gorm:"uniqueIndex"`
	Size string `json:"size"`
	Time string `json:"time"`
	Type string `json:"type"`

	// One-to-many relationship with PeerInfo
	Peers []PeerInfo `json:"peers" gorm:"foreignKey:FileMetadataID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
