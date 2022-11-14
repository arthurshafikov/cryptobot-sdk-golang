package cryptobot

const (
	testingAPIURL = "https://testnet-pay.crypt.bot/api/"
	realAPIURL    = "https://pay.crypt.bot/api/"
)

type Client struct {
	testingMode bool
	apiToken    string
}

type Deps struct {
	Testing  bool
	ApiToken string
}

func NewClient(deps Deps) *Client {
	return &Client{
		testingMode: deps.Testing,
		apiToken:    deps.ApiToken,
	}
}

func (c *Client) getRequestUrl() string {
	if c.testingMode {
		return testingAPIURL
	} else {
		return realAPIURL
	}
}
