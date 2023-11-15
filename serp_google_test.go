package brightdatasdk

import (
	"os"
	"testing"
)

var BDSerpClient *BrightDataClient

func TestMain(m *testing.M) {
	bdUsername := os.Getenv("BRIGHTDATA_USERNAME")
	if bdUsername == "" {
		panic("BRIGHTDATA_USERNAME is not set")
	}

	bdSerpPassword := os.Getenv("BRIGHTDATA_SERP_PASSWORD")
	if bdSerpPassword == "" {
		panic("BRIGHTDATA_SERP_PASSWORD is not set")
	}

	BDSerpClient = NewBrightDataClient(bdUsername, BrightDataCredentials{
		Serp: bdSerpPassword,
	})

	// Run the tests
	os.Exit(m.Run())
}

func TestSearchGoogleOutputJson(t *testing.T) {
	searchResult, err := BDSerpClient.GoogleSearch("brightdata", false, "en", "us")
	if err != nil {
		t.Fatalf("Error in GoogleSearch: %v", err)
	}

	if len(searchResult.Data.Organic) == 0 {
		t.Errorf("No search results returned")
	}
}

func TestSearchGoogleOutputHTML(t *testing.T) {
	searchResult, err := BDSerpClient.GoogleSearch("brightdata", true, "en", "us")
	if err != nil {
		t.Fatalf("Error in GoogleSearch: %v", err)
	}

	if len(searchResult.Html) == 0 {
		t.Errorf("No search results returned")
	}
}
