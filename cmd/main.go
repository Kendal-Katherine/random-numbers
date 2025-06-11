package main

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

func main(){
r := gin.Default()
r.GET("/random", func(c *gin.Context) {
	// Generate a random number and return it as JSON
	c.JSON(200, gin.H{
		"number": rand.Intn(100), // Random number between 0 and 99
	})
})

r.Run(":8080")

}