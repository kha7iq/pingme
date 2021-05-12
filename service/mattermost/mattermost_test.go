package mattermost

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockDoType
type MockDoType func(req *http.Request) (*http.Response, error)

// MockClient is the mock client
type MockClient struct {
	MockDo MockDoType
}

// Overriding what the Do function should "do" in our MockClient
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestSendMessage_Success(t *testing.T) {
	// build our response JSON
	successResponse, _ := json.Marshal(matterMostResponse{
		ID:      "1",
		Message: "Success",
	})
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader(successResponse))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	mattermostOpts := matterMost{
		Title:     "title",
		Token:     "token",
		ServerURL: "url",
		Scheme:    "https",
		APIURL:    "api-url",
		Message:   "hello",
		ChanIDs:   "1",
	}

	endPointURL := mattermostOpts.Scheme + "://" + mattermostOpts.ServerURL + mattermostOpts.APIURL

	// Create a Bearer string by appending string access token
	bearer := "Bearer " + mattermostOpts.Token

	fullMessage := mattermostOpts.Title + "\n" + mattermostOpts.Message

	ids := strings.Split(mattermostOpts.ChanIDs, ",")

	for _, v := range ids {
		assert.Equal(t, 1, len(v))

		jsonData, err := toJSON(v, fullMessage)
		assert.Nil(t, err)

		err = sendMattermost(endPointURL, bearer, jsonData)
		assert.Nil(t, err)
	}
}
