// SDK for Bright Data's proxy APIs implemented in GoLang
package brightdatasdk

type credentials struct {
	serp      string
	unblocker string
}

// BrightDataClient is the main client for the BrightData API
type BrightDataClient struct {
	customerID  string
	credentials credentials
}

// NewBrightDataClient creates a new BrightDataClient
func NewBrightDataClient(customerID string) *BrightDataClient {
	return &BrightDataClient{customerID: customerID}
}

// AuthenticateSerp authenticates the client for SERP requests
func (client *BrightDataClient) AuthenticateSerp(serpPassword string) *BrightDataClient {
	client.credentials.serp = serpPassword
	return client
}

// AuthenticateUnblocker authenticates the client for Unblocker requests
func (client *BrightDataClient) AuthenticateUnblocker(unblockerPassword string) *BrightDataClient {
	client.credentials.unblocker = unblockerPassword
	return client
}
