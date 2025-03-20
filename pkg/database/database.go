package database

import (
	"log"
	"sync"

	"github.com/Yogksai/rest-api-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Connect() *gorm.DB {
	once.Do(func() {
		cfg := config.NewConfig()
		dsn := "host=" + cfg.DB.Host + " port=" + cfg.DB.Port + " user=" + cfg.DB.User + " dbname=" + cfg.DB.Name + " password=" + cfg.DB.Password + " sslmode=disable"

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database")
		}
	})

	return db
}
