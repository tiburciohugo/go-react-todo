package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/tiburciohugo/go-react-todo/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	err := godotenv.Load()
	if err != nil {
		logger.Errorf("error loading .env file: %v", err)
		return nil, err
	}
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"))

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		logger.Errorf("error opening database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Todo{})
	if err != nil {
		logger.Errorf("error automigrating database: %v", err)
		return nil, err
	}
	return db, nil
}
