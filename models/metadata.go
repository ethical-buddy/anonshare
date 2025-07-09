package models

import (
	"gorm.io/gorm"
	"errors"
	
	
)

type FileMetadata struct {
	gorm.Model
	Hash string `json:"hash" gorm:"uniqueIndex"`
	Size string `json:"size"`
	Time string `json:"time"`
	Type string `json:"type"`

	// One-to-many relationship with PeerInfo
	Peers []PeerInfo `json:"peers" gorm:"foreignKey:FileMetadataID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}


// type Files


func UploadingInDB(db *gorm.DB, fileMetadata *FileMetadata) error {
	var existingFile FileMetadata
    result := db.Where("hash = ?", fileMetadata.Hash).First(&existingFile)

    peer := fileMetadata.Peers[0] // Only one peer per upload

    if result.Error == nil {
        // File exists, check if peer_id exists for this file
        var existingPeer PeerInfo
        peerResult := db.Where("file_metadata_id = ? AND peer_id = ?", existingFile.ID, peer.Peer_ID).First(&existingPeer)
        if peerResult.Error == gorm.ErrRecordNotFound {
            // Peer does not exist, add new peer
            peer.FileMetadataID = existingFile.ID
            if err := db.Create(&peer).Error; err != nil {
               return errors.New("failed to add peer to existing file")
            }
            return nil
        } else if peerResult.Error == nil {
            // Peer exists, update info
            existingPeer.IP = peer.IP
            existingPeer.Port = peer.Port
            existingPeer.FilePath = peer.FilePath
            existingPeer.Filename = peer.Filename
            existingPeer.Description = peer.Description
            if err := db.Save(&existingPeer).Error; err != nil {
				return errors.New("failed to update peer info")
            }
            
            
        } else {
           return errors.New("failed to check existing peer")
            
        }
    } else if result.Error == gorm.ErrRecordNotFound {
        // File does not exist, create new file with peer
        if err := db.Create(&fileMetadata).Error; err != nil {
           return errors.New("failed to create file metadata")
        }
       
        
    } else {
        return errors.New("database error")

    }
	return nil
}

func GettingFilesFromDB(db *gorm.DB,fileMetadata *[]FileMetadata) (error) {
	
	if err := db.Find(&fileMetadata).Error; err != nil {
		return errors.New("failed to fetch files from database")
	}
	return nil	
}


func GettingPeersFromDB(db *gorm.DB, hash string) ([]PeerMetadata, error) {
	var peers []PeerInfo
	err := db.Where("file_metadata_id IN (SELECT id FROM file_metadata WHERE hash = ?)", hash).Find(&peers).Error
	if err != nil {
		return nil, errors.New("failed to fetch peers from database")
	}

	var response []PeerMetadata
	for _, peer := range peers {
		response = append(response, PeerMetadata{
			Peer_ID:     peer.Peer_ID,
			IP:          peer.IP,
			Port:        peer.Port,
			FilePath:    peer.FilePath,
			Filename:    peer.Filename,
			Description: peer.Description,
		})
	}

	return response, nil
}

