package slack_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chanpon2015/slack-web-hook/slack"
)

func TestSend(t *testing.T) {
	webHook := slack.WebHook{
		URL: "",
		Msg: slack.Message{},
	}
	webHook.Msg.AddButton()
	data, err := json.Marshal(&webHook.Msg)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
	if err := webHook.Send(); err != nil {
		t.Fatal(err)
	}
}
