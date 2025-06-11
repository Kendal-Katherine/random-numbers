package handler

import "fmt"

type Request struct {
	Seed string `json:"seed" binding:"required"`
}

func (r *Request) Validate() error {
	if r.Seed == "" {
		return fmt.Errorf("seed is required")
	}

	return nil
}
