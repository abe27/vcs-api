package configs

import (
	"github.com/abe27/cvst20/api/models"
	"gorm.io/gorm"
)

var (
	StoreVCST       *gorm.DB
	StoreFormula    *gorm.DB
	StoreITC        *gorm.DB
	API_TRIGGER_URL string
	APP_NAME        string
	APP_DESCRIPTION string
)

func SetDB() {
	if !StoreFormula.Migrator().HasTable(&models.Employee{}) {
		StoreFormula.AutoMigrate(&models.Employee{})
	}

	if !StoreFormula.Migrator().HasTable(&models.Product{}) {
		StoreFormula.AutoMigrate(&models.Product{})
	}
}
