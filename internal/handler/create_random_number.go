package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"random-numbers/infrastructure/database"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateRandomNumber(ctx *gin.Context) {

	request := Request {}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		sendErrorResponse(ctx, http.StatusBadRequest, "Failed to parse request body")
		return
	}
	if err := request.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	input := database.RandomNumber{
		ID:        uuid.New().String(),
		Number:    rand.Intn(100),
		Seed:      request.Seed,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	
	if err := db.Create(&input).Error; err != nil {
		logger.Errorf("Failed to create random number: %v", err)
		sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create random number")
		return
	}

	logger.Infof("Request received with seed: %s", request)
	sendSuccessResponse(ctx, http.StatusCreated, input)
}
