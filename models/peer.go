package models

import (

)

type PeerInfo struct {
	Peer_ID uint   `json:"peer_id" gorm:"primaryKey"` // hash of the ip and port
	IP   string `json:"ip"`
	Port int    `json:"port"`
	FilePath  string `json:"file_path"`
	Filename  string `json:"filename"`
	Description string `json:"description"`
}