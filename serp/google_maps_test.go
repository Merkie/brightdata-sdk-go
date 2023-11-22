package serp

import (
	"testing"
)

func TestGoogleMaps(t *testing.T) {
	req, error := Serp.NewGoogleMapsRequest("pizza", "us", "en", 0, 10)
	if error != nil {
		panic(error)
	}

	resp, error := req.Execute()
	if error != nil {
		panic(error)
	}

	if len(resp.Organic) == 0 {
		t.Error("No results")
	}

	if resp.Organic[0].Title == "" {
		t.Error("No title for first result")
	}

	t.Logf("Found %d results, first result: %s\n", len(resp.Organic), resp.Organic[0].Title)
}
