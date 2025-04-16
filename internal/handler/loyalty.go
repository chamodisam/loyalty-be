package handler

import (
	"fmt"
	"loyalty-program-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Redeem points handler
func RedeemPoints(c *gin.Context) {
	var requestBody struct {
		AccountID string `json:"account_id"` // JSON field name should match the incoming JSON body
		Points    string `json:"points"`
	}

	// Bind the JSON request body to the struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Convert points from string to integer
	points, err := strconv.Atoi(requestBody.Points)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid points value",
		})
		return
	}
	accountID := requestBody.AccountID

	// Call the service function to redeem points
	err = service.RedeemPoints(accountID, points)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Redeemed %d loyalty points", points),
	})
}

// Earn points handler
func EarnPoints(c *gin.Context) {
	var requestBody struct {
		AccountID string `json:"account_id"` // JSON field name should match the incoming JSON body
		Points    string `json:"points"`
	}

	// Bind the JSON request body to the struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Convert points from string to integer
	points, err := strconv.Atoi(requestBody.Points)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid points value",
		})
		return
	}
	accountID := requestBody.AccountID

	// Call the service function to earn points
	err = service.EarnPoints(accountID, points)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Earned %d loyalty points", points),
	})
}

// Get the loyalty balance for a user
func GetBalance(c *gin.Context) {
	// Get account_id from query parameter
	accountID := c.DefaultQuery("account_id", "")

	// Call the service to retrieve the balance
	balance, err := service.GetBalance(accountID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return the balance
	c.JSON(http.StatusOK, gin.H{
		"balance": balance,
	})
}
