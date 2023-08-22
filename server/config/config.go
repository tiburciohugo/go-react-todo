package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Initialize() error {
	var err error

	// db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	// db, err = gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=go_react_todo port=5432 sslmode=disable"), &gorm.Config{})
	db, err = InitializeDB()

	if err != nil {
		return fmt.Errorf("error initializing PostgreSQL: %v", err)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
