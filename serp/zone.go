package serp

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/merkie/brightdata-sdk-go/util"
)

type SerpZone struct {
	client *http.Client
}

func NewSerpZone(BrdCustomerId, ZoneName, ZonePassword string) (*SerpZone, error) {
	// Build proxy URL as string
	proxyUrl := fmt.Sprintf("http://brd-customer-%s-zone-%s:%s@brd.superproxy.io:22225", BrdCustomerId, ZoneName, ZonePassword)

	// Parse the proxy URL
	parsedProxyUrl, err := url.Parse(proxyUrl)
	if err != nil {
		return nil, err
	}

	// Make the HTTP client
	client := util.CreateHttpClient(parsedProxyUrl)

	// Return all the things
	return &SerpZone{
		client: client,
	}, nil
}
