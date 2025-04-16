package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Google OAuth2 config
var GoogleOauthConfig = oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),                // Replace with your Google Client ID
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),            // Replace with your Google Client Secret
	RedirectURL:  "http://localhost:8080/auth/google/callback", // Google redirect URL
	Scopes:       []string{"email", "profile"},
	Endpoint:     google.Endpoint,
}

// Load environment variables and setup Google OAuth config
func init() {
	// Load environment variables from .env file
	err := godotenv.Load("/Users/chamodisamarawickrama/Documents/qlub-assignment/loyalty-program-be/.env")
	GoogleOauthConfig = oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),                // Google Client ID from the .env file
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),            // Google Client Secret from the .env file
		RedirectURL:  "http://localhost:8080/auth/google/callback", // Google Redirect URL
		Scopes:       []string{"email", "profile"},                 // Scopes for Google OAuth
		Endpoint:     google.Endpoint,                              // Google OAuth endpoint
	}
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Function to validate the Google OAuth token and get user info
func GetGoogleUserInfo(code string) (map[string]interface{}, error) {
	// Exchange the code for a token
	token, err := GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %v", err)
	}

	// Use the token to get user information
	client := GoogleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}

	return userInfo, nil
}

// Function to generate JWT token for authenticated user
func GenerateJWTToken(username string) (string, error) {
	// Define JWT claims (payload)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      jwt.TimeFunc().Add(time.Hour * 72).Unix(), // 72 hours expiration time
	}

	// Create the token using the claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}
