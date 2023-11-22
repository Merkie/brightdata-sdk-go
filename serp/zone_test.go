package serp

import (
	"os"
	"testing"
)

var Serp *SerpZone

func TestMain(M *testing.M) {
	// Create serp zone
	serp, err := NewSerpZone(os.Getenv("BRIGHTDATA_CUSTOMER_ID"), "serp", os.Getenv("BRIGHTDATA_SERP_PASSWORD"))
	if err != nil {
		panic(err)
	}
	Serp = serp

	// Run tests
	os.Exit(M.Run())
}
