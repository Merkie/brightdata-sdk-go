package unblocker

import (
	"net/http"

	"github.com/merkie/brightdata-sdk-go/util"
)

type UnblockerRequest struct {
	*UnblockerZone
	Url string
}

func (zone *UnblockerZone) NewRequest(Url string) *UnblockerRequest {
	return &UnblockerRequest{
		UnblockerZone: zone,
		Url:           Url,
	}
}

func (request *UnblockerRequest) Execute() (string, error) {
	// create request
	req, err := http.NewRequest("GET", request.Url, nil)
	if err != nil {
		return "", err
	}

	// add gzip header
	req.Header.Add("Accept-Encoding", "gzip")

	// send request
	resp, err := request.client.Do(req)
	if err != nil {
		return "", err
	}

	// Read response
	body, err := util.ReadGzipResponse(resp)
	if err != nil {
		return "", err
	}

	// Return body as string
	return string(body), nil
}
