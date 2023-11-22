package unblocker

import (
	"os"
	"testing"
)

var Unblocker *UnblockerZone

func TestMain(M *testing.M) {
	// Create unblocker zone
	ub, err := NewUnblockerZone(os.Getenv("BRIGHTDATA_CUSTOMER_ID"), "unblocker", os.Getenv("BRIGHTDATA_UNBLOCKER_PASSWORD"), "", "", "")
	if err != nil {
		panic(err)
	}
	Unblocker = ub

	// Run tests
	os.Exit(M.Run())
}
