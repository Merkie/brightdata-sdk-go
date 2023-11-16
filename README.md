# Bright Data SDK for Go

![Go Version 1.21.3](https://img.shields.io/badge/Go-1.21.3-blue)
[![Go Reference](https://pkg.go.dev/badge/github.com/merkie/brightdata-sdk-go.svg)](https://pkg.go.dev/github.com/merkie/brightdata-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/merkie/brightdata-sdk-go)](https://goreportcard.com/report/github.com/merkie/brightdata-sdk-go)

The Bright Data SDK for Go provides a convenient and efficient way to interact with [Bright Data](https://brightdata.com/)'s proxy APIs using GoLang.

> [!IMPORTANT]
> This project is currently in an **incomplete** state. Some expected features are yet to be implemented.

## Installation

Install the SDK using the following command:

```bash
go get -u "github.com/merkie/brightdata-sdk-go"
```

## Authentication

To use the Bright Data SDK, you need your customer ID and the passwords for the services you wish to use. These can be obtained from a Bright Data proxy URL, formatted as follows:

```
http://brd-<CUSTOMER ID>-zone-xxx:<SERVICE PASSWORD>@brd.superproxy.io:22225
```

## Usage Example

Set your Bright Data credentials in your environment:

**.bashrc** *(or your shell's equivalent configuration file)*
```bash
export BRIGHTDATA_CUSTOMER_ID=your_customer_id
export BRIGHTDATA_SERP_PASSWORD=your_serp_password
```

Example usage in a Go project:

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

	// Perform a Google search for "brightdata"
	searchResult, err := client.GoogleSearch("brightdata").CountryCode("us").Lang("en").Execute()
	if err != nil {
		panic(err)
	}

	// Print the first result as JSON
	jsonData, err := json.MarshalIndent(searchResult.Organic[0], "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))

	// You can also access the raw HTML too!
	fmt.Println(searchResult.Html[:200], "...")
}
```

## Roadmap

- [ ] Error handling guidelines
- [ ] Full implementation of all SERP functions
- [ ] Support for various proxy services (data center, ISP, unblocker, residential, mobile)
- [ ] Verbose logging capabilities
- [ ] Usage tracking (request count and cost)
- [ ] RAM-based and SQLite-based caching options
- [ ] Utility functions for asynchronous job endpoints
