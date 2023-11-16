package brightdatasdk

import (
	"testing"
)

func TestSearchGoogleMaps(t *testing.T) {
	mapsResponse, err := Client.GoogleMaps("hotels in Houston, tx").CountryCode("us").Lang("en").Pagination(10, 0).Execute()
	if err != nil {
		t.Error(err)
	}

	// Checks
	if err != nil {
		t.Fatalf("Error in GoogleMaps: %v", err)
	}

	if mapsResponse == nil {
		t.Fatalf("mapsResponse is nil")
	}

	if len(mapsResponse.Organic) == 0 {
		t.Fatalf("mapsResponse.Organic is empty")
	}

	if mapsResponse.Organic[0].Address == "" {
		t.Fatalf("mapsResponse.Organic[0].Address is empty")
	}

	if mapsResponse.Html == "" {
		t.Fatalf("mapsResponse.Html is empty")
	}
}
