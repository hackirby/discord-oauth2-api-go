package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"

	"discord-oauth2/internal/app/database/models"
)

var (
	Database *gorm.DB
)

func Connect() error {
	var err error
	Database, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_DSN")), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func Migrate() error {
	err := Database.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}
