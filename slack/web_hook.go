package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// WebHook is
type WebHook struct {
	URL string
	Msg Message
}

// Message is
type Message struct {
	// Text   string  `json:"text"`
	Blocks []Block `json:"blocks"`
	// Attachments []Attachment `json:"attachments"`
}

// AddButton is
func (m *Message) AddButton() Message {
	m.Blocks = append(m.Blocks, Block{
		Type: "actions",
		Elements: []Element{
			{
				Type: "button",
				Text: Text{
					Type: "plain_text",
					Text: "test",
				},
				Value: "test",
				URL:   "https://chanpon2015.com",
			},
		},
	})
	return *m
}

// Block is
type Block struct {
	BlockID  string    `json:"block_id"`
	Type     string    `json:"type"`
	Elements []Element `json:"elements"`
}

// Element is
type Element struct {
	Type  string `json:"type"`
	Text  Text   `json:"text"`
	Value string `json:"value"`
	URL   string `json:"url"`
}

// Text is
type Text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji"`
}

// Accessory is
type Accessory struct {
	Type  string `json:"type"`
	Text  Text   `json:"text"`
	Value string `json:"value"`
}

// Attachment is
type Attachment struct {
	Text    string    `json:"text"`
	Actions *[]Action `json:"actions"`
}

// Action is
type Action struct {
	Name  string `json:"name"`
	Text  string `json:"text"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

var httpClient http.Client

func init() {
	httpClient = *http.DefaultClient
}

// SetClient is
func SetClient(client http.Client) {
	httpClient = client
}

// Send is
func (w *WebHook) Send() error {
	msg, err := json.Marshal(w.Msg)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		http.MethodPost,
		w.URL,
		bytes.NewBuffer(msg),
	)
	if err != nil {
		return err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}

	return nil
}
