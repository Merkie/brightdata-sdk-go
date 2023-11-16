package brightdatasdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type googleMapsRequest struct {
	query       string
	lang        string
	countryCode string
	results     int
	skip        int
	client      *BrightDataClient
}

// GoogleMaps creates a new googleMapsRequest struct that can be executed via the Execute function
func (client *BrightDataClient) GoogleMaps(query string) *googleMapsRequest {
	return &googleMapsRequest{
		query:       query,
		lang:        "en",
		countryCode: "us",
		results:     10,
		skip:        0,
		client:      client,
	}
}

// Lang sets the language of the search
func (request *googleMapsRequest) Lang(lang string) *googleMapsRequest {
	request.lang = lang
	return request
}

// CountryCode sets the country of the search
func (request *googleMapsRequest) CountryCode(countryCode string) *googleMapsRequest {
	request.countryCode = countryCode
	return request
}

// Pagination sets the pagination of the search
func (request *googleMapsRequest) Pagination(results int, skip int) *googleMapsRequest {
	request.results = results
	request.skip = skip
	return request
}

// Execute executes the google maps request
func (request *googleMapsRequest) Execute() (*GoogleMapsResponse, error) {
	httpClient, err := request.client.getserpHTTPClient()
	if err != nil {
		return nil, err
	}

	// make the url
	url := fmt.Sprintf("https://www.google.com/maps/search/%s/?q=%s&gl=%s&lang=%s&start=%s&num=%s&brd_json=html", url.QueryEscape(request.query), url.QueryEscape(request.query), url.QueryEscape(request.countryCode), url.QueryEscape(request.lang), url.QueryEscape(fmt.Sprint(request.skip)), url.QueryEscape(fmt.Sprint(request.results)))

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

	var mapsResponse GoogleMapsResponse
	err = json.Unmarshal(body, &mapsResponse)
	if err != nil {
		return nil, err
	}

	return &mapsResponse, nil
}
