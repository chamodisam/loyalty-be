package main

import (
	"fmt"
	"loyalty-program-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	r := gin.Default()

	// Define the routes

	r.POST("/api/redeem", handler.RedeemPoints)

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
