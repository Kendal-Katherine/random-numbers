package config

import (
	gorm "gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {

	return nil
}

func GetLogger(p string) *Logger {
	// Initialize the logger
	logger := NewLogger(p)
	return logger
}
