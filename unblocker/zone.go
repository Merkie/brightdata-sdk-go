package unblocker

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/merkie/brightdata-sdk-go/util"
)

type UnblockerZone struct {
	client *http.Client
}

func NewUnblockerZone(BrdCustomerId, ZoneName, ZonePassword, Country, State, City string) (*UnblockerZone, error) {
	// Build proxy URL as string
	proxyUrl := fmt.Sprintf("http://brd-customer-%s-zone-%s", BrdCustomerId, ZoneName)

	if Country != "" {
		proxyUrl += fmt.Sprintf("-country-%s", Country)
	}

	if State != "" {
		proxyUrl += fmt.Sprintf("-state-%s", State)
	}

	if City != "" {
		proxyUrl += fmt.Sprintf("-city-%s", City)
	}

	proxyUrl += fmt.Sprintf(":%s@brd.superproxy.io:22225", ZonePassword)

	// Parse the proxy URL
	parsedProxyUrl, err := url.Parse(proxyUrl)
	if err != nil {
		return nil, err
	}

	// Make the HTTP client
	client := util.CreateHttpClient(parsedProxyUrl)

	// Return all the things
	return &UnblockerZone{
		client: client,
	}, nil
}
