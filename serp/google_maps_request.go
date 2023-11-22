package serp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/merkie/brightdata-sdk-go/util"
)

type GoogleMapsRequest struct {
	*SerpZone
	Url string
}

func (zone *SerpZone) NewGoogleMapsRequest(Query string, Country string, Language string, startFrom uint16, maxResults uint16) (*GoogleMapsRequest, error) {
	// Build URL
	url := fmt.Sprintf("https://www.google.com/maps/search/%s/?q=%s&brd_json=html", url.QueryEscape(Query), url.QueryEscape(Query))

	if Country != "" {
		url += fmt.Sprintf("&gl=%s", Country)
	}

	if Language != "" {
		url += fmt.Sprintf("&lang=%s", Language)
	}

	if startFrom > 0 {
		url += fmt.Sprintf("&start=%d", startFrom)
	}

	if maxResults > 0 {
		url += fmt.Sprintf("&num=%d", maxResults)
	}

	// Return all the things
	return &GoogleMapsRequest{
		SerpZone: zone,
		Url:      url,
	}, nil
}

func (request *GoogleMapsRequest) Execute() (*GoogleMapsResponse, error) {
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
	var response GoogleMapsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Return response
	return &response, nil
}
