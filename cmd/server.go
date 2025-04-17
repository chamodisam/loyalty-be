package main

import (
	"fmt"
	"loyalty-program-backend/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5175"},                   // Allow frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		AllowCredentials: true,                                                // Allow cookies and authentication headers
	}))

	// Define the routes

	r.POST("/api/earn", handler.EarnPoints)
	r.POST("/api/redeem", handler.RedeemPoints)
	r.GET("/api/balance", handler.GetBalance)

	// Google OAuth routes
	r.GET("/auth/google", handler.Login)             // Redirect to Google OAuth login
	r.GET("/auth/google/callback", handler.Callback) // Handle Google OAuth callback
	r.POST("/auth/exchange", handler.ExchangeCode)

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
