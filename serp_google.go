package brightdatasdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type BrightDataGoogleSearchResponse struct {
	Data *GoogleSearchResult
	Html string
}

func (client *BrightDataClient) GoogleSearch(query string, html bool, lang string, countryCode string) (*BrightDataGoogleSearchResponse, error) {
	httpClient, err := client.getserpHTTPClient()
	if err != nil {
		return nil, err
	}

	// init url for google search
	url := &url.URL{
		Scheme: "http",
		Host:   "google.com",
		Path:   "/search",
	}

	// add all the params
	params := url.Query()
	params.Add("q", query)
	params.Add("lang", lang)
	params.Add("gl", countryCode)
	if !html {
		params.Add("brd_json", "1")
	}

	// add the params to the url
	url.RawQuery = params.Encode()

	// perform the request
	resp, err := httpClient.Get(url.String())
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
		return nil, fmt.Errorf("invalid auth, check your credentials")
	}

	// unmarshal the response if not html
	if !html {
		var searchResult GoogleSearchResult
		err = json.Unmarshal(body, &searchResult)
		if err != nil {
			return nil, err
		}

		return &BrightDataGoogleSearchResponse{Data: &searchResult}, nil
	}

	// just return the html as a string if html is true
	return &BrightDataGoogleSearchResponse{Html: string(body)}, nil
}
