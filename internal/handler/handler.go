package handler

import (
	"random-numbers/infrastructure/config"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")

	db = config.GetSQLiteDB()
	if db == nil {
		logger.Errorf("Database connection is nil, cannot initialize handler")
		return
	}
}
