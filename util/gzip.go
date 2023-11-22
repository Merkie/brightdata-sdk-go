package util

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

func ReadGzipResponse(resp *http.Response) ([]byte, error) {
	var reader io.ReadCloser
	var err error

	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %v", err)
		}
		defer reader.Close()
	} else {
		reader = resp.Body
	}

	return io.ReadAll(reader)
}
