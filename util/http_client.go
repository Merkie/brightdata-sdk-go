package util

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

func CreateHttpClient(parsedProxyUrl *url.URL) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(parsedProxyUrl),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // This skips SSL certificate verification
			},
		},
	}
}
