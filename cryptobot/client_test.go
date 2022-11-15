package cryptobot

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

var testToken = "testToken"

func getTestClient(t *testing.T) *Client {
	t.Helper()

	return &Client{
		apiToken:    testToken,
		testingMode: true,
		httpClient:  &http.Client{},
	}
}

func TestGetRequestUrl(t *testing.T) {
	testnetClient := &Client{testingMode: true}
	mainnetClient := &Client{}

	require.Equal(t, testnetAPIURL, testnetClient.getRequestUrl())
	require.Equal(t, mainnetAPIURL, mainnetClient.getRequestUrl())
}
