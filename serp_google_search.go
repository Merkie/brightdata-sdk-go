package brightdatasdk

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	url, err := url.Parse("https://www.google.com/search")
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Add("q", request.query)
	q.Add("gl", request.countryCode)
	q.Add("lang", request.lang)
	q.Add("start", fmt.Sprint(request.page))
	q.Add("num", fmt.Sprint(request.resultsPerPage))
	q.Add("brd_json", "html")
	if request.useMobileDevice { // add mobile device if needed
		q.Add("brd_mobile", "1")
	}
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

	// execute request
	resp, err := serpHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// read response body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// invalid auth check (error from brightdata)
	if string(body) == "Invalid Auth" {
		return nil, fmt.Errorf("invalid auth")
	}

	// parse response body
	var googleSearchResponse GoogleSearchResponse
	err = json.Unmarshal(body, &googleSearchResponse)
	if err != nil {
		fmt.Println("attempting gzip decoding")
		// Attempt gzip decoding on error
		gzipReader, gzipErr := gzip.NewReader(io.NopCloser(bytes.NewReader(body)))
		if gzipErr != nil {
			return nil, fmt.Errorf("failed to decode response: %v, gzip error: %v", err, gzipErr)
		}
		defer gzipReader.Close()

		decompressedBody, gzipErr := io.ReadAll(gzipReader)
		if gzipErr != nil {
			return nil, fmt.Errorf("failed to decompress gzip: %v", gzipErr)
		}

		err = json.Unmarshal(decompressedBody, &googleSearchResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to decode decompressed response: %v", err)
		}
	}

	return &googleSearchResponse, nil
}
