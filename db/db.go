package db

import (
	"log"

	"github.com/dadadam/sono-backend/config"
	"github.com/dadadam/sono-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error

	c := config.GetConfig()

	dbConfig := &gorm.Config{}
	db, err = gorm.Open(postgres.Open(c.GetDbUrl()), dbConfig)
	if err != nil {
		log.Fatalln(err)
	}

	migrate()
}

func GetDB() *gorm.DB {
	return db
}

func migrate() {
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Book{})
}
