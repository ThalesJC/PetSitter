package database

import (
	"PetSitter/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstace struct {
	Db *gorm.DB
}

var Petsitter DbInstace

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var db *gorm.DB
	var err error
	maxRetries := 10

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Tentativa %d: Banco ainda não disponível, tentando novamente em 2s...", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Erro ao conectar ao banco após %d tentativas: %v", maxRetries, err)
	}

	log.Println("✅ Conectado ao banco na porta 5432!")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("📦 Executando Migrations...")
	db.AutoMigrate(&models.User{}, &models.Pet{})

	Petsitter = DbInstace{Db: db}
}
