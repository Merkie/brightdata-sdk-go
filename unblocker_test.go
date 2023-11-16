package brightdatasdk

import (
	"testing"
)

func TestUnblocker(t *testing.T) {
	unblockerResponse, err := Client.Unblocker("https://www.linkedin.com/in/archer-calder").Execute()
	if err != nil {
		t.Fatalf("Error in Unblocker: %v", err)
	}

	if unblockerResponse == "" {
		t.Fatalf("unblockerResponse is empty")
	}
}
