package handler

import (
	"fmt"
	"random-numbers/internal/random/mersenne"
	"testing"
	"time"
)

func TestGenerateValidNumberWithSeed(t *testing.T) {
	seed := "2025-06-11T15:00:00.000Z"
	parsedSeed, err := time.Parse(time.RFC3339Nano, seed)
	if err != nil {
		t.Fatalf("failed to parse seed: %v", err)
	}

	mt := mersenne.NewMT19937(uint32(parsedSeed.UnixNano()))
	number := int(mt.ExtractNumber() % 100)

	fmt.Printf("Generated number: %d\n", number)

	if number < 0 || number > 99 {
		t.Errorf("generated number %d is out of expected range [0,99]", number)
	}
}
