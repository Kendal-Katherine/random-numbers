package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"random-numbers/infrastructure/config"
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
		Number:    rand.Intn(5),
		Seed:      request.Seed,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	number, err := config.FindByNumber(db, input.Number)
	if err != nil {
		if err.Error() == "record not found" {
			logger.Infof("Number not found in database: %d", input.Number)
			sendSuccessResponse(ctx, http.StatusCreated, input)
			return
		}
		logger.Errorf("Failed to find number in database: %v", err)
		sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to find number in database")
		return
	}
	if number != nil {
		logger.Infof("Number already exists in database: %d", input.Number)
		sendErrorResponse(ctx, http.StatusConflict, "Number already exists in database")
		return
	}

	// err := rand.Seed(time.Now().UnixNano())
	// if err != nil {
	// 	logger.Errorf("Failed to seed random number generator: %v", err)
	// 	sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to seed random number generator")
	// 	return
	// }
	
	if err := db.Create(&input).Error; err != nil {
		logger.Errorf("Failed to create random number: %v", err)
		sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create random number")
		return
	}

	logger.Infof("Request received with seed: %s", request)
	sendSuccessResponse(ctx, http.StatusCreated, input)
}

