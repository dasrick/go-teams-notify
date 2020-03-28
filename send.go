package goteamsnotify

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// API - interface of MS Teams notify
type API interface {
	Send(webhookURL string, webhookMessage MessageCard) error
}

type teamsClient struct {
	httpClient *http.Client
}

// NewClient - create a brand new client for MS Teams notify
func NewClient() API {
	client := teamsClient{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	return &client
}

// Send - will post a notification to MS Teams incomingWebhookURL
func (c teamsClient) Send(webhookURL string, webhookMessage MessageCard) error {
	// validate url
	valid, err := isValidWebhookURL(webhookURL)
	if !valid {
		return err
	}
	// prepare message
	webhookMessageByte, _ := json.Marshal(webhookMessage)
	webhookMessageBuffer := bytes.NewBuffer(webhookMessageByte)

	// prepare request (error not possible)
	req, _ := http.NewRequest("POST", webhookURL, webhookMessageBuffer)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	// do the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode >= 299 {
		err = errors.New("error on notification: " + res.Status)
		log.Println(err)
		return err
	}

	return nil
}

// MessageCard - struct of message card
type MessageCard struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	ThemeColor string `json:"themeColor,omitempty"`
}

// NewMessageCard - create new empty message card
func NewMessageCard() MessageCard {
	return MessageCard{}
}

// helper --------------------------------------------------------------------------------------------------------------

func isValidWebhookURL(webhookURL string) (bool, error) {
	// basic URL check
	_, err := url.Parse(webhookURL)
	if err != nil {
		return false, err
	}
	// only pass MS teams webhook URLs
	switch {
	case strings.HasPrefix(webhookURL, "https://outlook.office.com/webhook/"):
	case strings.HasPrefix(webhookURL, "https://outlook.office365.com/webhook/"):
	default:
		err = errors.New("invalid ms teams webhook url")
		return false, err
	}
	return true, nil
}
