package brightdatasdk

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

var serpHTTPClient *http.Client

func (client *BrightDataClient) getserpHTTPClient() (*http.Client, error) {
	if serpHTTPClient != nil {
		return serpHTTPClient, nil
	}

	// configure proxy
	proxyURL, err := url.Parse(fmt.Sprintf("http://brd-customer-%s-zone-serp:%s@brd.superproxy.io:22225", client.customerID, client.credentials.serp))
	if err != nil {
		return nil, err
	}

	// configure HTTP client
	serpHTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // This skips SSL certificate verification
			},
		},
	}

	return serpHTTPClient, nil
}
