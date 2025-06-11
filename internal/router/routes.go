package router

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func initializeRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api/v1")

	apiGroup.GET("/random-number", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Random number generated successfully",
			"id":        uuid.New().String(),
			"number":    rand.Intn(100),
			"seed":      "2025-06-02T12:00:00.123Z",
			"createdAt": time.Now().Format(time.RFC3339),
		})
	})
}
