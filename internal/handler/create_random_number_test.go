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

func TestReproducibilityWithSameSeed(t *testing.T) {
	seed := "2025-06-11T15:00:00.000Z"
	parsedSeed, _ := time.Parse(time.RFC3339Nano, seed)
	seedUint := uint32(parsedSeed.UnixNano())

	mt1 := mersenne.NewMT19937(seedUint)
	mt2 := mersenne.NewMT19937(seedUint)

	n1 := mt1.ExtractNumber() % 100
	fmt.Printf("Generated number from mt1: %d\n", n1)
	n2 := mt2.ExtractNumber() % 100
	fmt.Printf("Generated number from mt2: %d\n", n2)

	if n1 != n2 {
		t.Errorf("expected same number for same seed, got %d and %d", n1, n2)
	}
}
