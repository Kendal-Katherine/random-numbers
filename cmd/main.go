package main

import (
	"random-numbers/infrastructure/config"
	"random-numbers/internal/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Configurations
	err := config.Init()
	if err == nil {
		logger.Errorf("Error initializing configurations: %v", err)
		return
	}

	// Initialize the router
	router.Initialize()
}
