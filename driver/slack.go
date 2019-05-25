package driver

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Slack struct {
	channel string
	webhook string
}

type body struct {
	Channel  string `json:"channel"`
	Text     string `json:"text"`
	Username string `json:"username"`
}

func NewSlack(c, w string) *Slack {
	return &Slack{
		channel: c,
		webhook: w,
	}
}

func (s *Slack) Send(message string) error {
	body := body{
		Channel:  s.channel,
		Text:     message,
		Username: "新聞君",
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return err
	}
	req, err := http.NewRequest("POST", s.webhook, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
