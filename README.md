# Bright Data SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/merkie/brightdata-sdk-go.svg)](https://pkg.go.dev/github.com/merkie/brightdata-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/merkie/brightdata-sdk-go)](https://goreportcard.com/report/github.com/merkie/brightdata-sdk-go)
![License](https://img.shields.io/badge/license-MIT-green)

The Bright Data SDK for Go provides a convenient type-safe way to interact with [Bright Data](https://brightdata.com/)'s proxy APIs.

> [!IMPORTANT]
> This project is currently in an **incomplete** state. Some expected features are yet to be implemented.

## Installation

```bash
go get -u "github.com/merkie/brightdata-sdk-go@latest"
```

## Authentication

To use the Bright Data SDK, you need your customer ID and the passwords for the services you wish to use. These can be obtained from a Bright Data proxy URL, formatted as follows:

```
http://brd-<CUSTOMER ID>-zone-xxx:<SERVICE PASSWORD>@brd.superproxy.io:22225
```

## Usage

<details>
<summary>Google Search</summary>

### Code:

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/merkie/brightdata-sdk-go/serp"
)

func main() {
	// Your BrightData credentials
	BrdCustomerID := "..."
	BrdSerpPassword := "..."

	// The name of your SERP zone ("serp" is the default)
	SerpZoneName := "serp"

	// Connect the SERP zone to the SDK
	// This will not create a new zone, only connect to an existing one
	Serp, err := serp.NewSerpZone(BrdCustomerID, SerpZoneName, BrdSerpPassword)
	if err != nil {
		panic(err)
	}

	// Create a new Google Search request
	req, err := Serp.NewGoogleSearchRequest("github", "us", "en", 0, 10)
	if err != nil {
		panic(err)
	}

	// Execute the request
	resp, err := req.Execute()
	if err != nil {
		panic(err)
	}

	// *optional* Print the response as JSON
	json, err := json.MarshalIndent(resp.Organic[0], "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}
```

### Output:

```json
{
	"link": "https://github.com/",
	"display_link": "https://github.com",
	"title": "GitHub: Let's build from here · GitHub",
	"description": "GitHub is where over 100 million developers shape the future of software, together. Contribute to the open source community, manage your Git repositories, ...",
	"extensions": [
		{
			"type": "site_link",
			"extended": true,
			"text": "Login",
			"link": "https://github.com/login",
			"rank": 1
		},
		{
			"type": "site_link",
			"extended": true,
			"text": "Explore GitHub",
			"link": "https://github.com/explore",
			"rank": 2
		},
		{
			"type": "site_link",
			"extended": true,
			"text": "Join GitHub",
			"link": "https://github.com/signup",
			"rank": 3
		},
		{
			"type": "site_link",
			"extended": true,
			"text": "GitHub Desktop",
			"link": "https://desktop.github.com/",
			"rank": 4
		}
	],
	"rank": 1,
	"global_rank": 1
}
```

</details>

<details>
<summary>Google Maps</summary>

### Code:

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/merkie/brightdata-sdk-go/serp"
)

func main() {
	// Your BrightData credentials
	BrdCustomerID := "..."
	BrdSerpPassword := "..."

	// The name of your SERP zone ("serp" is the default)
	SerpZoneName := "serp"

	// Connect the SERP zone to the SDK
	// This will not create a new zone, only connect to an existing one
	Serp, err := serp.NewSerpZone(BrdCustomerID, SerpZoneName, BrdSerpPassword)
	if err != nil {
		panic(err)
	}

	// Create a new Google Search request
	req, err := Serp.NewGoogleMapsRequest("the white house", "us", "en", 0, 10)
	if err != nil {
		panic(err)
	}

	// Execute the request
	resp, err := req.Execute()
	if err != nil {
		panic(err)
	}

	// *optional* Print the response as JSON
	json, err := json.MarshalIndent(resp.Organic[0], "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}
```

### Output:

```json
{
	"title": "The White House",
	"display_link": "whitehouse.gov",
	"link": "https://www.whitehouse.gov/",
	"address": "1600 Pennsylvania Avenue NW, Washington, DC 20500",
	"phone": "+12024561111",
	"category": [
		{
			"title": "Federal government office",
			"id": "federal_government_office"
		},
		{
			"title": "Government office",
			"id": "government_office"
		},
		{
			"title": "Historical place",
			"id": "historic_site"
		},
		{
			"title": "Historical landmark",
			"id": "historical_landmark"
		},
		{
			"title": "Tourist attraction",
			"id": "tourist_attraction"
		}
	],
	"tags": [
		{
			"group_id": "accessibility",
			"group_title": "Accessibility",
			"key_id": "/geo/type/establishment_poi/has_wheelchair_accessible_entrance",
			"value_title": "Has wheelchair accessible entrance"
		},
		{
			"group_id": "accessibility",
			"group_title": "Accessibility",
			"key_id": "/geo/type/establishment_poi/has_wheelchair_accessible_parking",
			"value_title": "Has wheelchair accessible parking lot"
		}
	],
	"summary": "Iconic home of America's president",
	"description": "Landmark, historic home \u0026 office of the United States president, with tours for visitors.",
	"rating": 4,
	"reviews_cnt": 4,
	"latitude": 38.8976763,
	"longitude": -77.0365298,
	"claimed": true,
	"fid": "0x89b7b7bcdecbb1df:0x715969d86d0b76bf",
	"map_id_encoded": "ChIJ37HL3ry3t4kRv3YLbdhpWXE",
	"map_id": "0x89b7b7bcdecbb1df:0x715969d86d0b76bf",
	"map_link": "https://www.google.com/maps/place/data=!3m1!4b1!4m2!3m1!1s0x89b7b7bcdecbb1df:0x715969d86d0b76bf",
	"original_image": "https://lh5.googleusercontent.com/p/AF1QipNfdRntXqqTYW5swoWU2U76NXsf_5-4kvwvwOxN=w408-h272-k-no",
	"image": "https://lh5.googleusercontent.com/p/AF1QipNfdRntXqqTYW5swoWU2U76NXsf_5-4kvwvwOxN=w138-h92-k-no",
	"thumbnail": "https://lh5.googleusercontent.com/p/AF1QipNfdRntXqqTYW5swoWU2U76NXsf_5-4kvwvwOxN=w129-h86-k-no",
	"icon": "",
	"image_url": "https://lh5.googleusercontent.com/p/AF1QipNfdRntXqqTYW5swoWU2U76NXsf_5-4kvwvwOxN=w138-h92-k-no",
	"rank": 1
}
```

</details>

<details>
<summary>Unblocker (aka Web Unlocker)</summary>

### Code:

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/merkie/brightdata-sdk-go/unblocker"
)

func main() {
	// Your BrightData credentials
	BrdCustomerID := "..."
	BrdSerpPassword := "..."

	// The name of your Unblocker zone ("unblocker" is the default)
	UnblockerZoneName := "unblocker"

	// Create a new Unblocker zone
	// This will not create a new zone, only connect to an existing one
	//
	// Last three arguments are Country, City and State, these need to be enabled
	// in your BrightData Unblocker dashboard before you can use them
	Unblocker, err := unblocker.NewUnblockerZone(BrdCustomerID, UnblockerZoneName, BrdUnblockerPassword, "", "", "")
	if err != nil {
		panic(err)
	}

	// Create and execute the request
	resp, err := Unblocker.NewRequest("https://www.reddit.com/r/github/").Execute()
	if err != nil {
		panic(err)
	}

	// *optional* Print all post titles
	for _, post := range strings.Split(resp, `slot="full-post-link"`) {
		postTitle := ""
		postHref := ""

		lines := strings.Split(strings.Split(post, "</a>")[0], "\n")[0:5]
		for _, line := range lines {
			if strings.Contains(line, "aria-label=") {
				postTitle = strings.Split(strings.Split(line, `aria-label="`)[1], `"`)[0]
			}
			if strings.Contains(line, "href=") {
				postHref = strings.Split(strings.Split(line, `href="`)[1], `"`)[0]
			}
		}

		if postTitle != "" && postHref != "" {
			fmt.Printf("%s\n%s\n\n", postTitle, postHref)
		}
	}
}
```

### Output:

```
Have or know of a project on Github looking for contributors? Feel free to drop them down to add to the wiki page!
/r/github/comments/c4kccq/have_or_know_of_a_project_on_github_looking_for/

401 when doing HEAD request to github.com
/r/github/comments/1816n3r/401_when_doing_head_request_to_githubcom/

🌟 GitHub Challenge: Improve QA Bots with GH Actions - Crypto Attack Wiki 🌟
/r/github/comments/1818z9p/github_challenge_improve_qa_bots_with_gh_actions/
```

</details>

## Contributing

Contributions to `brightdata-sdk-go` are welcome! Please refer to the project's issues page on GitHub for planned improvements and feature requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Legal Disclaimer

This project is an independent, community-driven effort and is not officially affiliated with, endorsed by, or in any way connected to Bright Data, Google, Microsoft (Bing), Yandex, or any of their subsidiaries or affiliates. The names Bright Data, Google, Bing, Yandex, and any related trademarks are the property of their respective owners and are used here for identification purposes only. This project is developed under the MIT License and is not associated with any official offerings from these entities.
