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
	if !Store.Migrator().HasTable(&models.Employee{}) {
		Store.AutoMigrate(&models.Employee{})
	}

	if !Store.Migrator().HasTable(&models.Product{}) {
		Store.AutoMigrate(&models.Product{})
	}
}
