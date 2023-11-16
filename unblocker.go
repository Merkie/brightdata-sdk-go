package brightdatasdk

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var unblockerHttpClient *http.Client

func (client *BrightDataClient) getUnblockerHTTPClient() (*http.Client, error) {
	if unblockerHttpClient != nil {
		return unblockerHttpClient, nil
	}

	// configure proxy
	proxyURL, err := url.Parse(fmt.Sprintf("http://brd-customer-%s-zone-unblocker:%s@brd.superproxy.io:22225", client.customerID, client.credentials.unblocker))
	if err != nil {
		return nil, err
	}

	// configure HTTP client
	unblockerHttpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // This skips SSL certificate verification
			},
		},
	}

	return unblockerHttpClient, nil
}

type unblockerRequest struct {
	url    string
	client *BrightDataClient
}

// Unblocker creates a new unblockerRequest struct that can be executed via the Execute function
func (client *BrightDataClient) Unblocker(url string) *unblockerRequest {
	return &unblockerRequest{
		url:    url,
		client: client,
	}
}

// Execute executes the unblocker request
func (request *unblockerRequest) Execute() (string, error) {
	unblockerClient, err := request.client.getUnblockerHTTPClient()
	if err != nil {
		panic(err)
	}

	// create request
	req, err := http.NewRequest("GET", request.url, nil)
	if err != nil {
		panic(err)
	}

	// send request
	resp, err := unblockerClient.Do(req)
	if err != nil {
		panic(err)
	}

	// read response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body), nil
}
