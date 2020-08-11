package telegrandma

import (
	"log"
	"testing"
)

var BotToken = "42:Oaua_fdk"
var ChatID = "987654321"

func Test_Bot_Getting_Last_Updates_with_success(t *testing.T) {
	bot, err := NewBot(BotToken)
	if err != nil {
		t.Errorf("Unexpected error: [%s]", err)
	}

	hTTPClient := &hTTPClientMock{}
	hTTPClient.setResponseHTTPStatusCode(200)
	hTTPClient.setExpectedURL("https://api.telegram.org/bot42:Oaua_fdk/getUpdates")
	hTTPClient.setResponseBody(responseBody)
	bot.hTTPClient = hTTPClient

	updates, err := bot.GetUpdates()
	if err != nil {
		t.Errorf("Unexpected error: [%s]", err)
	}

	log.Println("updates.Result", updates.Result)
}

func Test_Bot_Sending_Message_with_success(t *testing.T) {
	bot, err := NewBot(BotToken)
	if err != nil {
		t.Errorf("Unexpected error: [%s]", err)
	}

	hTTPClient := &hTTPClientMock{}
	hTTPClient.setResponseHTTPStatusCode(200)
	hTTPClient.setExpectedURL("https://api.telegram.org/bot42:Oaua_fdk/sendMessage?chat_id=987654321&parse_mode=html&text=What+happens+at+Nana%27s%E2%80%A6+stays+at+Nana%27s.")
	bot.hTTPClient = hTTPClient

	msg := "What happens at Nana's… stays at Nana's."
	if success, err := bot.SendMessage(ChatID, msg); !success {
		t.Errorf("Unexpected error: The message was not sent. Error: [%s]", err)
	}
}

func Test_Bot_Sending_Html_Message_with_success(t *testing.T) {
	bot, err := NewBot(BotToken)
	if err != nil {
		t.Errorf("Unexpected error: [%s]", err)
	}

	hTTPClient := &hTTPClientMock{}
	hTTPClient.setResponseHTTPStatusCode(200)
	hTTPClient.setExpectedURL("https://api.telegram.org/bot42:Oaua_fdk/sendMessage?chat_id=987654321&parse_mode=html&text=%3Cb%3EIf+nothing+is+going+well%2C+call+your+grandmother.%3C%2Fb%3E")
	bot.hTTPClient = hTTPClient

	msg := "<b>If nothing is going well, call your grandmother.</b>"
	if success, err := bot.SendHTML(ChatID, msg); !success {
		t.Errorf("Unexpected error: The message was not sent. Error: [%s]", err)
	}
}

func Test_Bot_With_No_Token(t *testing.T) {
	if _, err := NewBot(""); err == nil {
		t.Error("Unexpected error: Bots should not be initialized without a token")
	}
}

func Test_Bot_With_No_ChatID(t *testing.T) {
	bot, err := NewBot(BotToken)
	if err != nil {
		t.Errorf("Unexpected error: [%s]", err)
	}

	if success, _ := bot.SendHTML("", "This message must not be sent"); success {
		t.Errorf("Unexpected error: Bots should not return success when ChatID is null")
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
