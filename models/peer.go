package models

import (

)

type PeerInfo struct {
	Peer_ID string `json:"peer_id"`
	IP   string `json:"ip"`
	Port string    `json:"port"`
	FilePath  string `json:"file_path"`
	Filename  string `json:"filename"`
	Description string `json:"description"`
}