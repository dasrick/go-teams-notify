package go_teams_notify

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type message struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	ThemeColor string `json:"themeColor,omitempty"`
}

// Send - will post a notification to MS Teams incomingWebhookUrl
func Send(incomingWebhookUrl string, webhookMessage message) error {
	// ToDo validate url
	// needs to look like: https://outlook.office.com/webhook/xxx

	// validate notification
	webhookMessageByte, err := json.Marshal(webhookMessage)
	if err != nil {
		return err
	}
	// ToDo set defaults - e.g. ThemeColor

	// send notification
	webhookMessageBuffer := bytes.NewBuffer(webhookMessageByte)
	res, err := http.Post(incomingWebhookUrl, "application/json", webhookMessageBuffer)
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
