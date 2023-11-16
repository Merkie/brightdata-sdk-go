package brightdatasdk

import (
	"os"
	"testing"
)

var Client *BrightDataClient

func TestMain(m *testing.M) {
	// get env variables
	customerID := os.Getenv("BRIGHTDATA_CUSTOMER_ID")
	if customerID == "" {
		panic("BRIGHTDATA_CUSTOMER_ID is not set")
	}

	serpPassword := os.Getenv("BRIGHTDATA_SERP_PASSWORD")
	if serpPassword == "" {
		panic("BRIGHTDATA_SERP_PASSWORD is not set")
	}

	unblockerPassword := os.Getenv("BRIGHTDATA_UNBLOCKER_PASSWORD")
	if serpPassword == "" {
		panic("BRIGHTDATA_SERP_PASSWORD is not set")
	}

	// Create and authenticate client
	Client = NewBrightDataClient(customerID)
	Client.AuthenticateSerp(serpPassword)
	Client.AuthenticateUnblocker(unblockerPassword)

	// Run the tests
	os.Exit(m.Run())
}
