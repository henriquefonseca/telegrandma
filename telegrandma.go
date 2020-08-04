package telegrandma

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
)

const (
	URLBase             = "https://api.telegram.org/"
	SendMessageEndpoint = "sendMessage"
	GetUpdatesEndpoint  = "getUpdates"
)

type Bot struct {
	BotToken   string
	ChatID     string
	HttpClient Requester
}

func NewBot(token string) (*Bot, error) {
	if len(strings.TrimSpace(token)) < 1 {
		return &Bot{}, errors.New("Token cannot be empty")
	}

	return &Bot{BotToken: token}, nil
}

func (bot *Bot) GetUpdates() (*GetUpdatesResponse, error) {
	urlTarget := buildRootURLFrom(bot.BotToken) + "/" + GetUpdatesEndpoint

	if bot.HttpClient == nil {
		bot.HttpClient = new(HttpClient)
	}

	resp, err := bot.HttpClient.Get(urlTarget, map[string]string{})
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var response GetUpdatesResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Printf("Error deconding Telegram response: %v\n", err)
		return &GetUpdatesResponse{}, err
	}

	return &response, nil
}

//SendMessage: Send Telegram Messages using bot
func (bot *Bot) SendMessage(chatID, content string) (bool, error) {
	if len(strings.TrimSpace(chatID)) > 0 {
		bot.ChatID = chatID
	}

	if len(strings.TrimSpace(bot.ChatID)) < 1 {
		return false, errors.New("ChatID cannot be empty")
	}

	urlTarget := prepareUrlWith(bot.BotToken, bot.ChatID, content)

	if bot.HttpClient == nil {
		bot.HttpClient = new(HttpClient)
	}

	resp, err := bot.HttpClient.Get(urlTarget, nil)
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	log.Printf("Telegram response: %v\n", newStr)

	if err != nil {
		log.Printf("Request Error: [%v]\n", err)
		return false, err
	}

	return true, nil
}

//SendHTML: Send HTML through Telegram Messages using bot
func (bot *Bot) SendHTML(chatID, content string) (bool, error) {
	if len(strings.TrimSpace(chatID)) > 0 {
		bot.ChatID = chatID
	}

	if len(strings.TrimSpace(bot.ChatID)) < 1 {
		return false, errors.New("ChatID cannot be empty")
	}
	urlTarget := prepareUrlWith(bot.BotToken, bot.ChatID, content)

	if bot.HttpClient == nil {
		bot.HttpClient = new(HttpClient)
	}

	resp, err := bot.HttpClient.Get(urlTarget, nil)
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	log.Printf("Telegram response: %v\n", newStr)

	if err != nil {
		log.Printf("Request Error: [%v]\n", err)
		return false, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Printf("Request Error. Response Body: [%s]\n", resp.Body)
		return false, errors.New(fmt.Sprintf("Request Error: Status code [%d]\n", resp.StatusCode))
	}

	return true, nil
}

func prepareUrlWith(token, chatId, content string) string {
	baseUrl, err := url.Parse(URLBase)
	if err != nil {
		log.Println("Malformed URL: ", err.Error())
	}

	baseUrl.Path += "bot" + token + "/" + SendMessageEndpoint

	params := url.Values{}
	params.Add("chat_id", chatId)
	params.Add("text", content)
	params.Add("parse_mode", "html")

	baseUrl.RawQuery = params.Encode() // Escape Query Parameters

	return baseUrl.String()
}

func buildRootURLFrom(token string) string {
	return URLBase + "bot" + token
}
