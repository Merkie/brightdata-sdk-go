package brightdatasdk

import (
	"crypto/tls"
	"fmt"
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
		return "", err
	}

	// create request
	req, err := http.NewRequest("GET", request.url, nil)
	if err != nil {
		return "", err
	}

	// add gzip header
	req.Header.Add("Accept-Encoding", "gzip")

	// send request
	resp, err := unblockerClient.Do(req)
	if err != nil {
		return "", err
	}

	// Read response
	body, err := ReadResponse(resp)
	if err != nil {
		return "", err
	}

	// invalid auth check (error from brightdata)
	if string(body) == "Invalid Auth" {
		return "", fmt.Errorf("invalid auth")
	}

	return string(body), nil
}
