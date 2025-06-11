package router

import (
	"random-numbers/internal/handler"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	//Initialize Handlers
	handler.InitializeHandler()

	//Initialize Router
	r := gin.Default()

	//Initialize Routes
	initializeRoutes(r)
	//Run Server
	r.Run(":8080")
}
