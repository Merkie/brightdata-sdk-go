package brightdatasdk

type BrightDataCredentials struct {
	Serp string
}

type BrightDataClient struct {
	Username    string
	Credentials BrightDataCredentials
}

func NewBrightDataClient(username string, credentials BrightDataCredentials) *BrightDataClient {
	return &BrightDataClient{
		Username:    username,
		Credentials: credentials,
	}
}
