package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/matidev200/guardaloya-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}

var DB *gorm.DB

func NewDatabase() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	PORT := os.Getenv("DB_PORT") 

	URL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", USER, PASS, HOST, PORT, DBNAME)

	fmt.Println(URL)

	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&models.User{}, &models.Credential{})
	if err != nil {
		log.Fatalf("Error in migrate db: %v", err)
	}

	fmt.Println("Database connection established")

	DB = db
}
