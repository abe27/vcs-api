package configs

import (
	"github.com/abe27/cvst20/api/models"
	"gorm.io/gorm"
)

var (
	Store           *gorm.DB
	API_TRIGGER_URL string
	APP_NAME        string
	APP_DESCRIPTION string
)

func SetDB() {
	if !Store.Migrator().HasTable(&models.EMPLR{}) {
		Store.AutoMigrate(&models.EMPLR{})
	}
}
