package telegrandma

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var BotToken = "42:Oaua_fdk"
var ChatID = "987654321"

type HttpClientMock struct{}

func (hcMock *HttpClientMock) Get(url string, headers map[string]string) (*http.Response, error) {
	log.Printf("Mocking request to url %v\n", url)

	r := ioutil.NopCloser(bytes.NewReader([]byte(responseBody)))

	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil
}

func Test_Bot_Getting_Last_Updates(t *testing.T) {
	bot, err := NewBot(BotToken)
	if err != nil {
		t.Errorf("Unexpected error: [%d]\n", err)
	}

	bot.SetHttpClient(&HttpClientMock{})

	updates, err := bot.GetUpdates()
	if err != nil {
		t.Errorf("Unexpected error: [%d]\n", err)
	}

	log.Println("updates.Result", updates.Result)
}

func Test_Bot_Sending_Message(t *testing.T) {
	bot, err := NewBot(BotToken)
	if err != nil {
		t.Errorf("Unexpected error: [%d]\n", err)
	}

	bot.SetHttpClient(&HttpClientMock{})

	msg := "What happens at Nana's… stays at Nana's."
	success, err := bot.SendMessage(ChatID, msg)
	if !success {
		t.Errorf("Unexpected error: The message was not sent. Error: %s", err)
	}
}

func Test_Bot_Sending_Html_Message(t *testing.T) {
	bot, err := NewBot(BotToken)
	if err != nil {
		t.Errorf("Unexpected error: [%d]\n", err)
	}

	bot.SetHttpClient(&HttpClientMock{})

	msg := "<b>If nothing is going well, call your grandmother.</b>"
	success, err := bot.SendHTML(ChatID, msg)
	if !success {
		t.Errorf("Unexpected error: The message was not sent. Error: %s", err)
	}
}

var responseBody string = `{
	"ok": true,
	"result": [{
		"update_id": 47364,
		"message": {
			"message_id": 12334,
			"from": {
				"id": 375937,
				"is_bot": false,
				"first_name": "Jane",
				"last_name": "Marple",
				"username": "janemarple",
				"language_code": "en"
			},
			"chat": {
				"id": 375937,
				"first_name": "Jane",
				"last_name": "Marple",
				"username": "janemarple",
				"type": "private"
			},
			"date": 1596489691,
			"text": "People with a grudge against the world are always dangerous. They seem to think life owes them something"
		}
	}]
}`
