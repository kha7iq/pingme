package zulip_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/kha7iq/pingme/service/zulip"
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

	successResponse, _ := json.Marshal(zulip.ZResponse{
		ID:      1,
		Message: "",
		Result:  "success",
		Code:    "",
	})
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader(successResponse))

	zulip.Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	z := zulip.Zulip{
		ZBot: zulip.ZBot{
			EmailID: "test@test.com",
			APIKey:  "api-key",
		},
		Type:    "stream",
		To:      "general",
		Topic:   "test",
		Content: "test content",
		Domain:  "user.zulipchat.com",
	}

	resp, err := zulip.SendZulipMessage(z.Domain, z)

	assert.Nil(t, err)

	assert.Equal(t, "success", resp.Result)
}

func TestSendMessageStream_Fail(t *testing.T) {
	// build our response JSON

	failureResponse, _ := json.Marshal(zulip.ZResponse{
		Message: "Stream 'nonexistent_stream' does not exist",
		Result:  "error",
		Code:    "STREAM_DOES_NOT_EXIST",
	})
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader(failureResponse))

	zulip.Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body:       r,
			}, nil
		},
	}

	z := zulip.Zulip{
		ZBot: zulip.ZBot{
			EmailID: "test@test.com",
			APIKey:  "api-key",
		},
		Type:    "stream",
		To:      "general",
		Topic:   "test",
		Content: "test content",
		Domain:  "user.zulipchat.com",
	}

	resp, err := zulip.SendZulipMessage(z.Domain, z)

	assert.Nil(t, err)

	assert.Equal(t, "error", resp.Result)
}

func TestSendMessagePrivate_Fail(t *testing.T) {
	// build our response JSON

	failureResponse, _ := json.Marshal(zulip.ZResponse{
		Message: "some error",
		Result:  "error",
		Code:    "BAD_REQUEST",
	})
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader(failureResponse))

	zulip.Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body:       r,
			}, nil
		},
	}

	z := zulip.Zulip{
		ZBot: zulip.ZBot{
			EmailID: "test@test.com",
			APIKey:  "api-key",
		},
		Type:    "private",
		To:      "1,2",
		Topic:   "test",
		Content: "test content",
		Domain:  "user.zulipchat.com",
	}

	resp, err := zulip.SendZulipMessage(z.Domain, z)

	assert.Nil(t, err)

	assert.Equal(t, "error", resp.Result)
}

func TestSendMessagePrivate_Success(t *testing.T) {
	// build our response JSON
	successResponse, _ := json.Marshal(zulip.ZResponse{
		Message: "",
		Result:  "success",
		ID:      1,
	})
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader(successResponse))

	zulip.Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	z := zulip.Zulip{
		ZBot: zulip.ZBot{
			EmailID: "test@test.com",
			APIKey:  "api-key",
		},
		Type:    "private",
		To:      "1,2",
		Topic:   "test",
		Content: "test content",
		Domain:  "user.zulipchat.com",
	}

	resp, err := zulip.SendZulipMessage(z.Domain, z)

	assert.Nil(t, err)

	assert.Equal(t, "success", resp.Result)
}
