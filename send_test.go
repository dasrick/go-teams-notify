package goteamsnotify

import (
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func xTestSend(t *testing.T) {
//	var tests = []struct {
//		reqURL string
//		reqMsg message
//		error  error
//	}{
//		// success
//		{
//			reqURL: "https://outlook.office.com/webhook/a42444f3-d59e-4caf-b979-6df7919460b7@04ebf399-0553-42f8-9e20-599e669641dd/IncomingWebhook/6303c4b585b34d5986ad2e0fe0ddf3e2/2051545f-8c1d-4a1c-a2cd-1b90a24a0b99",
//			reqMsg: message{
//				Title:      "title",
//				Text:       "text **bold** and *italic* annd ***both*** ... some ~~strike through - doent work~~ but what happens if <br> * this<br> * is<br> * a<br> * list<br>",
//				ThemeColor: "ff00cc",
//			},
//			error: nil,
//		},
//	}
//
//	for _, test := range tests {
//		err := Send(test.reqURL, test.reqMsg)
//		assert.IsType(t, test.error, err)
//	}
//}

func TestSendFailInalidWebhookURL(t *testing.T) {
	emptyMessage := message{}
	var tests = []struct {
		reqURL string
		reqMsg message
		error  error
	}{
		{
			reqURL: "ht\ttp://",
			reqMsg: emptyMessage,
			error:  &url.Error{},
		},
		{
			reqURL: "",
			reqMsg: emptyMessage,
			error:  errors.New(""),
		},
		{
			reqURL: "http://foo.com/ctl\x80",
			reqMsg: emptyMessage,
			error:  errors.New(""),
		},
	}
	for _, test := range tests {
		err := Send(test.reqURL, test.reqMsg)
		assert.IsType(t, test.error, err)
	}
}

//func TestSendFailInvalidMessage(t *testing.T)  {
//	validURL := "https://outlook.office.com/webhook/1234567890"
//	var tests = []struct {
//		reqURL string
//		reqMsg interface{}
//		error  error
//	}{
//		{
//			reqURL: validURL,
//			reqMsg: make(chan int),
//			error:  errors.New(""),
//		},
//		//{
//		//	reqURL: validURL,
//		//	reqMsg: message{},
//		//	error:  errors.New(""),
//		//},
//	}
//	for _, test := range tests {
//		err := Send(test.reqURL, test.reqMsg)
//		assert.IsType(t, test.error, err)
//	}
//}
