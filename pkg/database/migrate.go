package database

import (
	"log"

	"github.com/Yogksai/rest-api-go/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	log.Default().Println("Migration has been processed")
	return nil
}
