package brightdatasdk

type BrightDataCredentials struct {
	serp string
}

type BrightDataClient struct {
	username    string
	credentials BrightDataCredentials
}

func NewBrightDataClient(username string, credentials BrightDataCredentials) *BrightDataClient {
	return &BrightDataClient{
		username,
		credentials,
	}
}
