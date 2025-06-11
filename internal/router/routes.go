package router

import (
	"net/http"
	"random-numbers/infrastructure/config"
	"random-numbers/internal/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api/v1")
	apiGroup.Use(APIKeyAuthMiddleware())
	apiGroup.POST("/random-number", handler.CreateRandomNumber)
}

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-API-Key")
		expectedAPIKey := config.GetAPIKey()

		if apiKey == "" || apiKey != expectedAPIKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or missing API key",
			})
			return
		}

		ctx.Next()
	}
}
