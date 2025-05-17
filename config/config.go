package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Debug     bool   `json:"debug"`
	DataDir   string `json:"data_dir"`
	ListenPort int    `json:"listen_port"`
	PeerPort   int    `json:"peer_port"`
}

const defaultConfigPath = "anonshare_config.json"

func LoadConfig() (*Config, error) {
	if _, err := os.Stat(defaultConfigPath); os.IsNotExist(err) {
		fmt.Println("Config file not found. Creating default config...")
		cfg := defaultConfig()
		err := saveConfig(cfg)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}

	file, err := os.ReadFile(defaultConfigPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func defaultConfig() *Config {
	return &Config{
		Debug:     true,
		DataDir:   filepath.Join(".", "data"),
		ListenPort: 8080,
		PeerPort:   7070,
	}
}

func saveConfig(cfg *Config) error {
	bytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(defaultConfigPath, bytes, 0644)
}

