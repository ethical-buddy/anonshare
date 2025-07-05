package database

import (
	"log"

	"github.com/muskiteer/anonshare/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("metadata.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.FileMetadata{})
	if err != nil {
		log.Fatal("failed to migrate:", err)
	}

	return db
}
