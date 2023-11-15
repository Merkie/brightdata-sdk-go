package brightdatasdk

type BrightDataCredentials struct {
	serp string
}

type BrightDataClient struct {
	username    string
	credentials BrightDataCredentials
}

func NewBrightDataClient(username string) *BrightDataClient {
	return &BrightDataClient{username: username}
}

func (client *BrightDataClient) AuthenticateSerp(serpPassword string) {
	client.credentials.serp = serpPassword
}
