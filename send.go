package goteamsnotify

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// MessageCard - struct of message card
type MessageCard struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	ThemeColor string `json:"themeColor,omitempty"`
}

// Send - will post a notification to MS Teams incomingWebhookURL
func Send(incomingWebhookURL string, webhookMessage MessageCard) error {
	// validate url
	// needs to look like: https://outlook.office.com/webhook/xxx
	valid, err := isValidWebhookURL(incomingWebhookURL)
	if valid != true {
		return err
	}

	// ToDo set defaults - e.g. ThemeColor

	// send notification
	webhookMessageByte, _ := json.Marshal(webhookMessage)
	webhookMessageBuffer := bytes.NewBuffer(webhookMessageByte)
	res, err := http.Post(incomingWebhookURL, "application/json", webhookMessageBuffer)
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

// NewMessageCard - create new empty message card
func NewMessageCard() MessageCard {
	return MessageCard{}
}

func isValidWebhookURL(webhookURL string) (bool, error) {
	// basic URL check
	_, err := url.Parse(webhookURL)
	if err != nil {
		return false, err
	}
	// only pass MS teams webhook URLs
	hasPrefix := strings.HasPrefix(webhookURL, "https://outlook.office.com/webhook/")
	if hasPrefix != true {
		err = errors.New("unvalid ms teams webhook url")
		return false, err
	}
	return true, nil
}
