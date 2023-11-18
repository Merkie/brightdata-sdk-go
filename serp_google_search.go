package brightdatasdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type googleSearchRequest struct {
	query           string
	lang            string
	countryCode     string
	resultsPerPage  int
	page            int
	useMobileDevice bool
	client          *BrightDataClient
}

// GoogleSearch creates a new googleSearchRequest struct that can be executed via the Execute function
func (client *BrightDataClient) GoogleSearch(query string) *googleSearchRequest {
	return &googleSearchRequest{
		query:           query,
		lang:            "en",
		countryCode:     "us",
		resultsPerPage:  10,
		page:            0,
		useMobileDevice: false,
		client:          client,
	}
}

// Lang sets the language of the search
func (request *googleSearchRequest) Lang(lang string) *googleSearchRequest {
	request.lang = lang
	return request
}

// CountryCode sets the country of the search
func (request *googleSearchRequest) CountryCode(countryCode string) *googleSearchRequest {
	request.countryCode = countryCode
	return request
}

// Pagination sets the pagination of the search
func (request *googleSearchRequest) Pagination(resultsPerPage int, page int) *googleSearchRequest {
	request.resultsPerPage = resultsPerPage
	request.page = page
	return request
}

// UseMobileDevice enables mobile device emulation
func (request *googleSearchRequest) UseMobileDevice(useMobileDevice bool) *googleSearchRequest {
	request.useMobileDevice = useMobileDevice
	return request
}

// Execute executes the google search request
func (request *googleSearchRequest) Execute() (*GoogleSearchResponse, error) {
	httpClient, err := request.client.getserpHTTPClient()
	if err != nil {
		return nil, err
	}

	// make the url
	url := fmt.Sprintf("https://www.google.com/search?q=%s&gl=%s&lang=%s&start=%s&num=%s&brd_json=1", url.QueryEscape(request.query), url.QueryEscape(request.countryCode), url.QueryEscape(request.lang), url.QueryEscape(fmt.Sprint(request.page)), url.QueryEscape(fmt.Sprint(request.resultsPerPage)))

	// add mobile device if needed
	if request.useMobileDevice {
		url = url + "&brd_mobile=1"
	}

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
