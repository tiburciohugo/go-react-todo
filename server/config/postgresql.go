package config

import (
	"os"

	"github.com/tiburciohugo/go-react-todo/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("database file does not exist, creating it: %s", dbPath)
		// create the database file
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("error creating database directory: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("error creating database file: %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(postgres.Open(dbPath), &gorm.Config{})
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
