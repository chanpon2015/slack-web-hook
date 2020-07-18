package slack_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chanpon2015/slack-web-hook/slack"
)

func TestSend(t *testing.T) {
	webHook := slack.WebHook{
		URL: "https://hooks.slack.com/services/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",
		Msg: slack.Message{},
	}
	actions := webHook.Msg.CreateActions()
	actions.AddButton("test1", "", "").
		AddButton("test2", "https://chanpon2015.com", "")
	data, err := json.Marshal(&webHook.Msg)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
	if err := webHook.Send(); err != nil {
		t.Fatal(err)
	}
}
