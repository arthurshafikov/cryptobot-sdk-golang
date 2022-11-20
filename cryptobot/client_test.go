package cryptobot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
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

func TestRequest(t *testing.T) {
	c := getTestClient(t)
	bodyContent := "123"
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/somePath").
		MatchHeader(apiTokenHeaderName, testToken).
		Reply(200).
		BodyString(bodyContent)

	responseBodyReader, err := c.request("somePath", nil)
	require.NoError(t, err)

	body, err := ioutil.ReadAll(responseBodyReader)
	require.NoError(t, err)
	require.Equal(t, bodyContent, string(body))
}

func TestRequestModifiedResponse(t *testing.T) {
	c := getTestClient(t)
	defer gock.Off()
	gock.New(testnetAPIURL).
		Get("/somePath").
		MatchHeader(apiTokenHeaderName, testToken).
		MatchParam("param1", "someValue1").
		MatchParam("param2", "someValue2").
		Reply(200).
		BodyString(`123`)

	responseBodyReader, err := c.request("somePath", func(q url.Values) url.Values {
		q.Add("param1", "someValue1")
		q.Add("param2", "someValue2")

		return q
	})
	require.NoError(t, err)

	body, err := ioutil.ReadAll(responseBodyReader)
	require.NoError(t, err)
	require.Equal(t, "123", string(body))
}

func TestDecodeResponse(t *testing.T) {
	c := getTestClient(t)
	expectedTransfer := Transfer{
		ID:          123,
		UserID:      "1234",
		Asset:       USDT,
		Amount:      "200",
		Status:      "paid",
		CompletedAt: "",
		Comment:     "someComment",
	}
	transferJSON, err := json.Marshal(expectedTransfer)
	require.NoError(t, err)

	var target Transfer
	err = c.decodeResponse(bytes.NewReader(transferJSON), &target)
	require.NoError(t, err)
	require.Equal(t, expectedTransfer, target)
}

func TestDecodeResponseWrongJSON(t *testing.T) {
	c := getTestClient(t)

	var target Transfer
	err := c.decodeResponse(bytes.NewReader([]byte("wrong json...")), &target)
	require.ErrorContains(
		t,
		err,
		"error while unmarshaling response to the target: invalid character 'w' looking for beginning of value",
	)
}
