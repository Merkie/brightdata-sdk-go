package brightdatasdk

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

var SerpHTTPClient *http.Client

func (client *BrightDataClient) getSerpHTTPClient() (*http.Client, error) {
	if SerpHTTPClient != nil {
		return SerpHTTPClient, nil
	}

	// configure proxy
	proxyURL, err := url.Parse(fmt.Sprintf("http://%s-zone-serp:%s@brd.superproxy.io:22225", client.Username, client.Credentials.Serp))
	if err != nil {
		return nil, err
	}

	// configure HTTP client
	SerpHTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // This skips SSL certificate verification
			},
		},
	}

	return SerpHTTPClient, nil
}
