package cryptobot

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	testingAPIURL      = "https://testnet-pay.crypt.bot/api/"
	realAPIURL         = "https://pay.crypt.bot/api/"
	apiTokenHeaderName = "Crypto-Pay-API-Token"
)

type Client struct {
	testingMode bool
	apiToken    string

	httpClient *http.Client
}

type Deps struct {
	Testing  bool
	ApiToken string

	// Optional. Default is 10 seconds
	ClientTimeout time.Duration
}

func NewClient(deps Deps) *Client {
	c := &Client{
		testingMode: deps.Testing,
		apiToken:    deps.ApiToken,
	}

	clientTimeout := time.Second * 10
	if deps.ClientTimeout != 0 {
		clientTimeout = deps.ClientTimeout
	}
	c.httpClient = &http.Client{
		Timeout: clientTimeout,
	}

	return c
}

func (c *Client) getRequestUrl() string {
	if c.testingMode {
		return testingAPIURL
	} else {
		return realAPIURL
	}
}

func (c *Client) request(path string, queryModifierFunc func(q url.Values) url.Values) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, c.getRequestUrl()+path, nil)
	if err != nil {
		return nil, err
	}
	if queryModifierFunc != nil {
		req.URL.RawQuery = queryModifierFunc(req.URL.Query()).Encode()
	}

	req.Header.Set(apiTokenHeaderName, c.apiToken)
	r, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return r.Body, nil
}
