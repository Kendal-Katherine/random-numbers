package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//Initialize Router
	r := gin.Default()

	//Initialize Routes
	initializeRoutes(r)
	//Run Server
	r.Run(":8080")
}
