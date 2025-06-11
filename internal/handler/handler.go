package handler

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateRandomNumber(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Random number generated successfully",
		"id":        uuid.New().String(),
		"number":    rand.Intn(100),
		"seed":      "2025-06-02T12:00:00.123Z",
		"createdAt": time.Now().Format(time.RFC3339),
	})
}
