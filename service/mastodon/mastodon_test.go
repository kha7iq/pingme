package mastodon

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const url, token = "server-url", "token"

// MockClient is the mock client
type MockClient struct {
	MockDo func(req *http.Request) (*http.Response, error)
}

// Do function implements HTTPClient
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestSendMessage_Success(t *testing.T) {
	successResponse, _ := json.Marshal(map[string]interface{}{
		"success": true,
	})

	r := ioutil.NopCloser(bytes.NewReader(successResponse))

	Client = &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, url, req.URL.Path)
			assert.Equal(t, token, req.Header.Get("Authorization"))

			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	m := Mastodon{
		ServerURL: url,
		Token:     token,
		Message:   "message",
	}

	err := sendMastodon(m.ServerURL, m.Token, m.Message)
	assert.Nil(t, err)
}

func TestSendMessage_Failure(t *testing.T) {
	successResponse, _ := json.Marshal(map[string]interface{}{
		"error": true,
	})

	r := ioutil.NopCloser(bytes.NewReader(successResponse))

	Client = &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, url, req.URL.Path)
			assert.Equal(t, token, req.Header.Get("Authorization"))

			return &http.Response{
				StatusCode: 400,
				Body:       r,
			}, nil
		},
	}

	m := Mastodon{
		ServerURL: url,
		Token:     token,
		Message:   "message",
	}

	err := sendMastodon(m.ServerURL, m.Token, m.Message)
	assert.NotNil(t, err)
}
