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

	// Create and authenticate client
	Client = NewBrightDataClient(customerID)
	Client.AuthenticateSerp(serpPassword)

	// Run the tests
	os.Exit(m.Run())
}

func TestSearchGoogle(t *testing.T) {
	// Perform the search
	searchResult, err := Client.GoogleSearch("brightdata", "en", "us")

	// Checks
	if err != nil {
		t.Fatalf("Error in GoogleSearch: %v", err)
	}

	if searchResult == nil {
		t.Fatalf("searchResult is nil")
	}

	if len(searchResult.Organic) == 0 {
		t.Fatalf("searchResult.Organic is empty")
	}

	if searchResult.Organic[0].Link == "" {
		t.Fatalf("searchResult.Organic[0].Link is empty")
	}

	if searchResult.Html == "" {
		t.Fatalf("searchResult.Html is empty")
	}
}
