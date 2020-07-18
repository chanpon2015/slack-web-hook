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
	URL    string
	Msg    Message
	Blocks []Block
}

// Message is
type Message struct {
	Msg    string  `json:"message,omitempty"`
	Blocks []Block `json:"blocks"`
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
