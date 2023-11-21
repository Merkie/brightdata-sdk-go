package brightdatasdk

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	url, err := url.Parse("https://www.google.com/maps/search/" + url.QueryEscape(request.query))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Add("q", request.query)
	q.Add("gl", request.countryCode)
	q.Add("lang", request.lang)
	q.Add("start", fmt.Sprint(request.skip))
	q.Add("num", fmt.Sprint(request.results))
	q.Add("brd_json", "html")
	url.RawQuery = q.Encode()

	// Get serp http client
	serpHTTPClient, err := request.client.getserpHTTPClient()
	if err != nil {
		return nil, err
	}

	// create request
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	// add gzip header
	req.Header.Add("Accept-Encoding", "gzip")

	// execute request
	resp, err := serpHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// print all headers
	for k, v := range resp.Header {
		fmt.Printf("%s: %s\n", k, v)
	} // one says Content-Encoding: [gzip]

	// Read response
	body, err := ReadResponse(resp)
	if err != nil {
		return nil, err
	}

	// invalid auth check (error from brightdata)
	if string(body) == "Invalid Auth" {
		return nil, fmt.Errorf("invalid auth")
	}

	// parse response body
	var mapsResponse GoogleMapsResponse
	err = json.Unmarshal(body, &mapsResponse)
	if err != nil {
		return nil, err
	}

	return &mapsResponse, nil
}
