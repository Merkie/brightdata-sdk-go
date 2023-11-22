package unblocker

import (
	"fmt"
	"strings"
	"testing"
)

func TestRequest(T *testing.T) {
	// Create request
	request := Unblocker.NewRequest("https://www.linkedin.com/in/williamhgates")

	// Execute request
	body, err := request.Execute()
	if err != nil {
		panic(err)
	}

	if len(body) == 0 {
		T.Errorf("Expected response body to be non-empty")
	}

	if !strings.Contains(body, `<script type="application/ld+json">`) {
		T.Errorf(`Expected response body to contain '<script type="application/ld+json">'`)
	}

	fmt.Printf("[Unblocker] Response body length: %d\n", len(body))
}
