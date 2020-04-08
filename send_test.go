package goteamsnotify

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	assert.IsType(t, &teamsClient{}, client)
}

func TestTeamsClientSend(t *testing.T) {
	// THX@Hassansin ... http://hassansin.github.io/Unit-Testing-http-client-in-Go
	simpleMsgCard := NewMessageCard()
	simpleMsgCard.Text = "Hello World"
	var tests = []struct {
		reqURL    string
		reqMsg    MessageCard
		resStatus int   // httpClient response status
		resError  error // httpClient error
		error     error // method error
	}{
		// invalid webhookURL - url.Parse error
		{
			reqURL:    "ht\ttp://",
			reqMsg:    simpleMsgCard,
			resStatus: 0,
			resError:  nil,
			error:     &url.Error{},
		},
		// invalid webhookURL - missing prefix in webhook URL
		{
			reqURL:    "",
			reqMsg:    simpleMsgCard,
			resStatus: 0,
			resError:  nil,
			error:     errors.New(""),
		},
		// invalid httpClient.Do call
		{
			reqURL:    "https://outlook.office.com/webhook/xxx",
			reqMsg:    simpleMsgCard,
			resStatus: 200,
			resError:  errors.New("pling"),
			error:     &url.Error{},
		},
		// invalid httpClient.Do call
		{
			reqURL:    "https://outlook.office365.com/webhook/xxx",
			reqMsg:    simpleMsgCard,
			resStatus: 200,
			resError:  errors.New("pling"),
			error:     &url.Error{},
		},
		// invalid response status code
		{
			reqURL:    "https://outlook.office.com/webhook/xxx",
			reqMsg:    simpleMsgCard,
			resStatus: 400,
			resError:  nil,
			error:     errors.New(""),
		},
		// invalid response status code
		{
			reqURL:    "https://outlook.office365.com/webhook/xxx",
			reqMsg:    simpleMsgCard,
			resStatus: 400,
			resError:  nil,
			error:     errors.New(""),
		},
		// valid
		{
			reqURL:    "https://outlook.office.com/webhook/xxx",
			reqMsg:    simpleMsgCard,
			resStatus: 200,
			resError:  nil,
			error:     nil,
		},
		// valid
		{
			reqURL:    "https://outlook.office365.com/webhook/xxx",
			reqMsg:    simpleMsgCard,
			resStatus: 200,
			resError:  nil,
			error:     nil,
		},
	}
	for _, test := range tests {
		client := NewTestClient(func(req *http.Request) (*http.Response, error) {
			// Test request parameters
			assert.Equal(t, req.URL.String(), test.reqURL)
			return &http.Response{
				StatusCode: test.resStatus,
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}, test.resError
		})
		c := &teamsClient{httpClient: client}

		err := c.Send(test.reqURL, test.reqMsg)
		assert.IsType(t, test.error, err)
	}
}

// helper for testing --------------------------------------------------------------------------------------------------

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) (*http.Response, error)

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// NewTestClient returns *http.API with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
