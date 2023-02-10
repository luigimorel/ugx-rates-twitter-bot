package main

import (
	"os"
	"testing"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	// Load .env file
	err := godotenv.Load()
	require.NoError(t, err, "Error loading .env file")

	// Get environment variables
	apiKey := os.Getenv("API_KEY")
	apiKeySecret := os.Getenv("API_KEY_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	// Test environment variables are not empty
	require.NotEmpty(t, apiKey, "API_KEY environment variable is empty")
	require.NotEmpty(t, apiKeySecret, "API_KEY_SECRET environment variable is empty")
	require.NotEmpty(t, accessToken, "ACCESS_TOKEN environment variable is empty")
	require.NotEmpty(t, accessTokenSecret, "ACCESS_TOKEN_SECRET environment variable is empty")

	config := oauth1.NewConfig(apiKey, apiKeySecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Test the VerifyCredentials method
	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	require.NoError(t, err, "Error verifying Twitter credentials")
	require.NotEmpty(t, user.ScreenName, "Twitter account screen name is empty")

	// Test the Update status method
	_, _, err = client.Statuses.Update(ConvertMoney(), nil)
	require.NoError(t, err, "Error updating Twitter status")
}
