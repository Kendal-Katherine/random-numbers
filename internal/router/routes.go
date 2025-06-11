package router

import (
	"random-numbers/internal/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api/v1")

	apiGroup.POST("/random-number", handler.CreateRandomNumber)
}
