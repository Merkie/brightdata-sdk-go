# Bright Data SDK for Go

![Go verion 1.21.3](https://img.shields.io/badge/Go-1.21.3-blue)
![0 Dependencies](https://img.shields.io/badge/Dependencies-0-blue)

SDK for [Bright Data](https://brightdata.com/)'s proxy APIs implemented in GoLang

> [!IMPORTANT]
> This project is currently functionally **incomplete**; features you may expect might still need to be implemented.

## Install

```bash
go get -u "github.com/merkie/brightdata-sdk-go"
```

## Authenticating

The two things needed to authenticate with Bright Data are your customer ID and the passwords of the services you plan on using with the SDK. To get these, copy one of the proxy URLs provided by Bright Data and match the example here: `http://brd-<CUSTOMER ID>-zone-xxx:<SERVICE PASSWORD>@brd.superproxy.io:22225`. If you want to use multiple services, all you will need is the passwords; you only submit your customer ID once.

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
	client := NewBrightDataClient(os.Getenv("BRIGHTDATA_CUSTOMER_ID"))
	client.AuthenticateSerp(os.Getenv("BRIGHTDATA_SERP_PASSWORD"))
	// client.AuthenticateDataCenter(...)
	// client.AuthenticateISP(...)
	// client.AuthenticateUnblocker(...)
	// ...

	query := "brightdata"

	searchResult, err := client.GoogleSearch(query, "en", "us")
	if err != nil {
		panic(err)
	}

	// At this point it's a fully-typed struct, let's unmarshal the first result and print it

	json, err := json.Marshal(searchResult.Data.Organic[0])
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
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
