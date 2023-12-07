package db

import (
	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/models"
)

func AutomigrateUserManagement(db *gorm.DB) error {
	var err error

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}

func AutomigrateWallet(db *gorm.DB) error {
	var err error
	err = db.AutoMigrate(&models.Transaction{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Credentials{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Wallet{})
	if err != nil {
		return err
	}

	return nil
}
