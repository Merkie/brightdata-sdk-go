package brightdatasdk

type BrightDataCredentials struct {
	serp string
}

type BrightDataClient struct {
	customerID  string
	credentials BrightDataCredentials
}

func NewBrightDataClient(customerID string) *BrightDataClient {
	return &BrightDataClient{customerID: customerID}
}

func (client *BrightDataClient) AuthenticateSerp(serpPassword string) {
	client.credentials.serp = serpPassword
}
