package brightdatasdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

func (client *BrightDataClient) GoogleSearch(query string, lang string, countryCode string) (*GoogleSearchResponse, error) {
	httpClient, err := client.getserpHTTPClient()
	if err != nil {
		return nil, err
	}

	// make the url
	url := fmt.Sprintf("https://google.com/search?brd_json=html&lang=%s&gl=%s&q=%s", url.QueryEscape(lang), url.QueryEscape(countryCode), url.QueryEscape(query))

	// perform the request
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response as bytes
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// invalid auth check (error from brightdata)
	if string(body) == "Invalid Auth" {
		return nil, fmt.Errorf("invalid auth")
	}

	var searchResult GoogleSearchResponse
	err = json.Unmarshal(body, &searchResult)
	if err != nil {
		return nil, err
	}

	return &searchResult, nil
}
