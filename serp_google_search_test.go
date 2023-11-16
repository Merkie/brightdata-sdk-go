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
	// searchResult, err := Client.GoogleSearch("brightdata", "en", "us")

	searchResponse, err := Client.GoogleSearch("brightdata").CountryCode("us").Lang("en").Pagination(10, 0).Execute()

	// Checks
	if err != nil {
		t.Fatalf("Error in GoogleSearch: %v", err)
	}

	if searchResponse == nil {
		t.Fatalf("searchResponse is nil")
	}

	if len(searchResponse.Organic) == 0 {
		t.Fatalf("searchResponse.Organic is empty")
	}

	if searchResponse.Organic[0].Link == "" {
		t.Fatalf("searchResponse.Organic[0].Link is empty")
	}

	if searchResponse.Html == "" {
		t.Fatalf("searchResponse.Html is empty")
	}
}
