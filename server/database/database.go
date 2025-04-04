package database

import (
	"PetSitter/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstace struct {
	Db *gorm.DB
}

var Petsitter DbInstace

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=petsitter port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error: ", err.Error())
		os.Exit(2)
	}

	log.Println("Connected at: 'port:5432'")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations...")

	db.AutoMigrate(&models.User{}, &models.Pet{})

	Petsitter = DbInstace{Db: db}
}
