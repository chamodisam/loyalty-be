package service

import (
	"context"
	"fmt"
	"log"

	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/square/square-go-sdk"
	"github.com/square/square-go-sdk/client"
	"github.com/square/square-go-sdk/loyalty"
	"github.com/square/square-go-sdk/option"
)

// Function to generate a unique idempotency key
func generateIdempotencyKey() string {
	// Generate a new UUID
	idempotencyKey := uuid.New().String()

	// Return the idempotency key
	return idempotencyKey
}

// Initialize Square Client
func getSquareClient() *client.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the Square Access Token from the environment variable
	token := os.Getenv("SQUARE_ACCESS_TOKEN")
	client := client.NewClient(
		option.WithBaseURL(
			square.Environments.Sandbox, // or use square.Environments.Production
		),
		option.WithToken(
			token, // Replace with your actual Square access token
		),
	)
	return client
}

// Redeem points for a user
func RedeemPoints(accountID string, points int) error {
	// Initialize Square client
	client := getSquareClient()
	idempotencyKey := generateIdempotencyKey()

	// Prepare the request to adjust (redeem) loyalty points
	redeemRequest := &loyalty.AdjustLoyaltyPointsRequest{
		AccountID: accountID,
		AdjustPoints: &square.LoyaltyEventAdjustPoints{
			Points: points,                                    // Adjust points to redeem
			Reason: square.String("Redeem points for reward"), // Reason for adjustment
		},
		IdempotencyKey: idempotencyKey,
	}

	// Call the Square API to redeem the points
	_, err := client.Loyalty.Accounts.Adjust(
		context.TODO(),
		redeemRequest,
	)

	if err != nil {
		log.Printf("Error redeeming points: %v", err)
		return fmt.Errorf("failed to redeem points for account %s", accountID)
	}

	return nil
}

// Get the loyalty account balance
func GetBalance(accountID string) (int, error) {
	// Initialize Square client
	client := getSquareClient()

	// Make the request to retrieve the loyalty account details
	response, err := client.Loyalty.Accounts.Get(
		context.TODO(),
		&loyalty.GetAccountsRequest{
			AccountID: accountID,
		},
	)

	if err != nil {
		log.Printf("Error retrieving loyalty account: %v", err)
		return 0, fmt.Errorf("failed to retrieve loyalty account for account %s", accountID)
	}

	// Return the loyalty account balance
	if response.LoyaltyAccount.Balance != nil {
		return *response.LoyaltyAccount.Balance, nil // Dereference the pointer
	}

	return 0, fmt.Errorf("balance is not available for account %s", accountID)
}
