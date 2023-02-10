package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConvertMoney(t *testing.T) {
	// Save the original values of MONEY_API_KEY and MONEY_API_BASE_URL
	originalMoneyApiKey := os.Getenv("MONEY_API_KEY")
	originalMoneyApiBaseUrl := os.Getenv("MONEY_API_BASE_URL")
	// Define the expected result from the API call
	expectedResult := CurrencyConvert{
		Result: 3590.0,
	}

	// Start a test server to simulate the API call
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the apikey header is present in the request
		apikey := r.Header.Get("apikey")
		require.Equal(t, "test_api_key", apikey)

		// Write the expected result as JSON
		err := json.NewEncoder(w).Encode(expectedResult)
		require.NoError(t, err)
	}))
	defer ts.Close()

	// Set the MONEY_API_KEY and MONEY_API_BASE_URL environment variables for the test
	os.Setenv("MONEY_API_KEY", "test_api_key")
	os.Setenv("MONEY_API_BASE_URL", ts.URL)

	// Call the ConvertMoney function
	result := ConvertMoney()

	// Check if the result matches the expected result
	expectedDate := time.Now().Format("02 January 2006, 15:04:05 PM") + "\n"
	expectedVersion := LogVersion()
	expectedFinalString := expectedDate + "\n💰 1 USD ............ UGX 3590\n\n" + expectedVersion
	require.Equal(t, expectedFinalString, result)

	// Reset the MONEY_API_KEY and MONEY_API_BASE_URL environment variables to their original values
	os.Setenv("MONEY_API_KEY", originalMoneyApiKey)
	os.Setenv("MONEY_API_BASE_URL", originalMoneyApiBaseUrl)
}
