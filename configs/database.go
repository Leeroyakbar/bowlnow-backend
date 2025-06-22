package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/Leeroyakbar/bowlnow-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB = db
	fmt.Println("Connected to database")

	DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Category{},
		&models.Food{},
		&models.TransactionDetail{},
		&models.PaymentMethod{},
		&models.Transaction{},
	)

}
