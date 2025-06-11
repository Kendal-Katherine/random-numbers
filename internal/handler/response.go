package handler

import (
	"random-numbers/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func sendErrorResponse(c *gin.Context, code int, message string) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}
func sendSuccessResponse(c *gin.Context, status int, data database.RandomNumber) {
	c.JSON(status, gin.H{
		"message":   "Random number generated successfully",
		"data":      data,
	})
}
