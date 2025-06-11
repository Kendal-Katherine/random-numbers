package config

import (
	"errors"
	"os"
	"random-numbers/infrastructure/database"

	"gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := ".db/main.db"

	//Check is the datavase file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Database file does not exist, creating directory and file: %s", dbPath)

		//Create the directory if it does not exist
		err = os.MkdirAll(".db", os.ModePerm)
		if err != nil {
			logger.Errorf("Failed to create database directory: %v", err)
			return nil, err
		}

		//Create the database file
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Failed to create database file: %v", err)
			return nil, err
		}
		defer file.Close()
		logger.Infof("Database file created successfully: %s", dbPath)
	}

	// Initialize the SQLite database connection
	logger.Infof("Initializing SQLite database connection...")
	db, err := gorm.Open(sqlite.Open(".db/main.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect to the database: %v", err)
		return nil, err
	}
	// Migrate the schema for the RandomNumber model
	err = db.AutoMigrate(&database.RandomNumber{})
	if err != nil {
		logger.Errorf("Failed to migrate database schema: %v", err)
		return nil, err
	}
	//Return the database connection
	logger.Infof("SQLite database connection initialized successfully.")
	return db, nil
}

func FindByNumber(db *gorm.DB, number int) (*database.RandomNumber, error) {
	var result database.RandomNumber

	err := db.Where("number = ?", number).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Retorna nil sem erro se não encontrado
		}
		return nil, err // Retorna erro se for erro real
	}

	return &result, nil
}
