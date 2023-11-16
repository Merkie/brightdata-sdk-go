package brightdatasdk

import (
	"testing"
)

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
