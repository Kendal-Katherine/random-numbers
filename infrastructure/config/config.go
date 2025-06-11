package config

import (
	"os"

	"github.com/joho/godotenv"
	gorm "gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	// Initialize the SQLite database connection
	db, err = InitializeSQLite()
	if err != nil {

		return err
	}
	return nil
}

func GetSQLiteDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize the logger
	logger := NewLogger(p)
	return logger
}


func LoadEnv() {
	logger := GetLogger("loadEnv")
	err := godotenv.Load("../.env")
	if err != nil {
		logger.Errorf("No .env file found or error loading .env")
	}
}

func GetAPIKey() string {
	return os.Getenv("API_KEY")
}