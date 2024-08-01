package database

import (
	"fmt"
	"log"
	"os"

	"github.com/sabrihanane/go-network-api-fiber-postgres/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connected")
}

func AutoMigrate() {

	DB.AutoMigrate(&models.Network{}, &models.Subnet{}, &models.Link{}, &models.Ltp{}, &models.Node{})
	//DB.AutoMigrate(&models.Network{}, &models.Subnet{}, &models.Link{}, &models.Node{})
}
