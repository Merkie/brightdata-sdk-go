package serp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/merkie/brightdata-sdk-go/util"
)

type GoogleSearchRequest struct {
	*SerpZone
	Url string
}

func (zone *SerpZone) NewGoogleSearchRequest(Query string, Country string, Language string, StartPage uint16, ResultsPerPage uint16) (*GoogleSearchRequest, error) {
	// Build URL
	url := fmt.Sprintf("https://www.google.com/search?q=%s&brd_json=html", url.QueryEscape(Query))

	if Country != "" {
		url += fmt.Sprintf("&gl=%s", Country)
	}

	if Language != "" {
		url += fmt.Sprintf("&lang=%s", Language)
	}

	if StartPage > 0 {
		url += fmt.Sprintf("&start=%d", StartPage)
	}

	if ResultsPerPage > 0 {
		url += fmt.Sprintf("&num=%d", ResultsPerPage)
	}

	// Return all the things
	return &GoogleSearchRequest{
		SerpZone: zone,
		Url:      url,
	}, nil
}

func (request *GoogleSearchRequest) Execute() (*GoogleSearchResponse, error) {
	// create request
	req, err := http.NewRequest("GET", request.Url, nil)
	if err != nil {
		return nil, err
	}

	// add gzip header
	req.Header.Add("Accept-Encoding", "gzip")

	// execute request
	resp, err := request.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Read response
	body, err := util.ReadGzipResponse(resp)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	var response GoogleSearchResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Return response
	return &response, nil
}
