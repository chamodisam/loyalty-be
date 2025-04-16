package handler

import (
	"loyalty-program-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// Login function to redirect to Google OAuth login
func Login(c *gin.Context) {

	url := service.GoogleOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

// Callback function to handle the response from Google OAuth
func Callback(c *gin.Context) {
	// Get code from query string
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code is missing"})
		return
	}

	// Call the service to get user info from Google using the code
	userInfo, err := service.GetGoogleUserInfo(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Optionally, you can generate JWT token here
	token, err := service.GenerateJWTToken(userInfo["email"].(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT token"})
		return
	}

	// Return user info and token
	c.JSON(http.StatusOK, gin.H{
		"user_info": userInfo,
		"token":     token,
	})
}
