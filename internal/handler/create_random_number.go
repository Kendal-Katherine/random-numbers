package handler

import (
	"encoding/json"
	"net/http"
	"random-numbers/infrastructure/config"
	"random-numbers/infrastructure/database"
	"random-numbers/internal/random/mersenne"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateRandomNumber(ctx *gin.Context) {

	request := Request{}

	//Try to decode the request body into the Request struct
	//It's ideal to validate seed 
	_ = json.NewDecoder(ctx.Request.Body).Decode(&request)

	// Check if the seed is provided in the request
	var seedTime time.Time
	if request.Seed == "" {
		seedTime = time.Now()
		request.Seed = seedTime.Format(time.RFC3339Nano)
	} else {
		parsedSeed, err := time.Parse(time.RFC3339Nano, request.Seed)
		if err != nil {
			sendErrorResponse(ctx, http.StatusBadRequest, "Invalid seed format (expected RFC3339Nano)")
			return
		}
		seedTime = parsedSeed
	}

	seed := uint32(seedTime.UnixNano())
	mt := mersenne.NewMT19937(seed)
	randomNumber := int(mt.ExtractNumber() % 100)

	input := database.RandomNumber{
		ID:        uuid.New().String(),
		Number:    randomNumber,
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

	if err := db.Create(&input).Error; err != nil {
		logger.Errorf("Failed to create random number: %v", err)
		sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create random number")
		return
	}

	logger.Infof("Request received with seed: %s", request)
	sendSuccessResponse(ctx, http.StatusCreated, input)
}
