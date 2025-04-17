package handler

import (
	"encoding/json"
	"loyalty-program-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// Login function to redirect to Google OAuth login
func Login(c *gin.Context) {

	url := service.GoogleOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	c.JSON(http.StatusOK, gin.H{"redirectUrl": url})
}

func ExchangeCode(c *gin.Context) {
	// Capture the authorization code sent by the frontend
	var req struct {
		Code string `json:"code"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Exchange code for token
	token, err := service.GoogleOauthConfig.Exchange(c, req.Code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to exchange code"})
		return
	}

	// Use the token to get the user's profile information
	client := service.GoogleOauthConfig.Client(c, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user info"})
		return
	}
	defer resp.Body.Close()

	// Parse user info (e.g., email, name)
	var userInfo struct {
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
		return
	}

	// Respond with the access token and user info to the frontend
	c.JSON(http.StatusOK, gin.H{
		"token": token.AccessToken, // Send the token back to the frontend
		"user":  userInfo,          // Send user info (email, name, picture)
	})
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
