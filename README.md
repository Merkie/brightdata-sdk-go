# Bright Data SDK for Go

![Go verion 1.21.3](https://img.shields.io/badge/Go-1.21.3-blue)
![75% Coverage](https://img.shields.io/badge/Test_Coverage-75%25-yellow)
[![Go Reference](https://pkg.go.dev/badge/github.com/merkie/brightdata-sdk-go.svg)](https://pkg.go.dev/github.com/merkie/brightdata-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/merkie/brightdata-sdk-go)](https://goreportcard.com/report/github.com/merkie/brightdata-sdk-go)
![0 Dependencies](https://img.shields.io/badge/Dependencies-0-blue)

SDK for [Bright Data](https://brightdata.com/)'s proxy APIs implemented in GoLang

> [!IMPORTANT]
> This project is currently functionally **incomplete**; features you may expect might still need to be implemented.

## Install

```bash
go get -u "github.com/merkie/brightdata-sdk-go"
```

## Authenticating

To authenticate with Bright Data, you need your customer ID and the passwords for the services you intend to use with the SDK. Obtain these by copying a proxy URL from Bright Data, formatted as `http://brd-<CUSTOMER ID>-zone-xxx:<SERVICE PASSWORD>@brd.superproxy.io:22225`. For multiple services, only the service passwords are required after the initial submission of your customer ID.

## Code Example

**.bashrc** *(or your shell's source file)*
```bash
export BRIGHTDATA_CUSTOMER_ID=...
export BRIGHTDATA_SERP_PASSWORD=...
```

**your-go-project/main.go**
```go
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
```

## Coming Soon...
* API.md
* Error Handling docs
* SERP docs
* Full support for all SERP functions
* Support for data center, ISP, unblocker, residential, and mobile services
* Verbose logging
* Usage tracking in request count and money spent
* RAM-based and SQLite-based caching
* Functions to help with making an async job endpoint
