package main

import (
	"encoding/json"
	"fmt"
	"os"

	brightdatasdk "github.com/merkie/brightdata-sdk-go"
)

func main() {
	client := brightdatasdk.NewBrightDataClient(os.Getenv("BRIGHTDATA_CUSTOMER_ID"))
	client.AuthenticateSerp(os.Getenv("BRIGHTDATA_SERP_PASSWORD"))
	// client.AuthenticateDataCenter(...)
	// client.AuthenticateISP(...)
	// client.AuthenticateUnblocker(...)
	// ...

	// Now that we are authenticated, let's perform a basic Google search for "brightdata"

	query := "brightdata"

	searchResult, err := client.GoogleSearch(query, "en", "us")
	if err != nil {
		panic(err)
	}

	// At this point it's a fully-typed struct, let's unmarshal the first result and print it

	json, err := json.MarshalIndent(searchResult.Organic[0], "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))

	// You can also access the raw html of the page as a string

	fmt.Println(searchResult.Html[:200], "...")
}
