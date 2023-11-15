# Bright Data SDK for Go

![Go verion 1.21.3](https://img.shields.io/badge/Go-1.21.3-blue)
![0 Dependencies](https://img.shields.io/badge/Dependencies-0-blue)

SDK for [Bright Data](https://brightdata.com/)'s proxy APIs implemented in GoLang

> [!IMPORTANT]
> This project is currently functionally **incomplete**, features you may expect might not be implemented yet.

## Install

```bash
go get -u "github.com/merkie/brightdata-sdk-go"
```

## Authenticating

There are two things you need to connect to the Bright Data SDK, you need your Bright Data username and the password(s) to the services you will be using.
As an example, if your proxy URL for SERP is `http://brd-customer-hl_userid-zone-serp:password123@brd.superproxy.io:22225`, you will insert `brd-customer-hl_userid` as your username and `password123` as your SERP authentication.

## Code Example

**.bashrc** *(or your shell's source file)*
```bash
export BRIGHTDATA_USERNAME=...
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
	client := brightdatasdk.NewBrightDataClient(os.Getenv("BRIGHTDATA_USERNAME"))
	client.AuthenticateSerp(os.Getenv("BRIGHTDATA_SERP_PASSWORD"))
	// client.AuthenticateDataCenter(...)
	// client.AuthenticateISP(...)
	// client.AuthenticateUnblocker(...)
	// ...

	searchResult, err := client.GoogleSearch("brightdata", false, "en", "us")
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
