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
	// URLBase is the Telegram base url
	URLBase = "https://api.telegram.org/"
	// SendMessageEndpoint is the Telegram sendMessage endpoint
	SendMessageEndpoint = "sendMessage"
	// GetUpdatesEndpoint is the Telegram getUpdates endpoint
	GetUpdatesEndpoint = "getUpdates"
)

// Bot is the Struct representing a telegram bot
type Bot struct {
	botToken   string
	chatID     string
	hTTPClient requester
}

// NewBot function creates a telegram bot with a token
//
// It returns an error if a token is not provided
func NewBot(token string) (*Bot, error) {
	if len(strings.TrimSpace(token)) < 1 {
		return &Bot{}, errors.New("token cannot be empty")
	}

	return &Bot{botToken: token}, nil
}

// GetUpdates returns an array of incoming updates from bot.
func (bot *Bot) GetUpdates() (*GetUpdatesResponse, error) {
	urlTarget := buildRootURLFrom(bot.botToken) + "/" + GetUpdatesEndpoint

	if bot.hTTPClient == nil {
		bot.hTTPClient = new(hTTPClient)
	}

	resp, err := bot.hTTPClient.Get(urlTarget, map[string]string{})
	defer resp.Body.Close()

	if err != nil {
		log.Printf("Error trying request for url: %s\n", urlTarget)
		return &GetUpdatesResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Error trying to read response body from url: %s\n", urlTarget)
		return &GetUpdatesResponse{}, err
	}

	var response GetUpdatesResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Printf("Error deconding Telegram response: %v\n", err)
		return &GetUpdatesResponse{}, err
	}

	return &response, nil
}

// SendMessage is used to send a message to chatID using bot
//
// It returns a boolean indicating if the operation was successful.
// If it is not, the error must provide a description about the problem.
func (bot *Bot) SendMessage(chatID, content string) (bool, error) {
	if len(strings.TrimSpace(chatID)) > 0 {
		bot.chatID = chatID
	}

	if len(strings.TrimSpace(bot.chatID)) < 1 {
		return false, errors.New("chatID cannot be empty")
	}

	urlTarget := prepareURLWith(bot.botToken, bot.chatID, content, "html")

	if bot.hTTPClient == nil {
		bot.hTTPClient = new(hTTPClient)
	}

	resp, err := bot.hTTPClient.Get(urlTarget, nil)
	defer resp.Body.Close()

	if err != nil {
		log.Printf("Request Error: [%v]\n", err)
		return false, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	log.Printf("Telegram response: %v\n", newStr)

	return true, nil
}

// SendHTML is used to send HTML to chatID using bot.
// Consult this url in order to check allowed tags:
// https://core.telegram.org/bots/api#html-style
//
// It returns a boolean indicating if the operation was successful.
// If it is not, the error must provide a description about the problem.
func (bot *Bot) SendHTML(chatID, content string) (bool, error) {
	if len(strings.TrimSpace(chatID)) > 0 {
		bot.chatID = chatID
	}

	if len(strings.TrimSpace(bot.chatID)) < 1 {
		return false, errors.New("chatID cannot be empty")
	}
	urlTarget := prepareURLWith(bot.botToken, bot.chatID, content, "html")

	if bot.hTTPClient == nil {
		bot.hTTPClient = new(hTTPClient)
	}

	resp, err := bot.hTTPClient.Get(urlTarget, nil)
	defer resp.Body.Close()

	if err != nil {
		log.Printf("Request Error: [%v]\n", err)
		return false, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	log.Printf("Telegram response: %v\n", newStr)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Printf("Request Error. Response Body: [%s]\n", resp.Body)
		return false, fmt.Errorf("Request error: Status code [%d]", resp.StatusCode)
	}

	return true, nil
}

// SendMarkdown is used to send markdown to chatID using bot.
// Consult this url in order to check allowed tags:
// https://core.telegram.org/bots/api#markdown-style
//
// It returns a boolean indicating if the operation was successful.
// If it is not, the error must provide a description about the problem.
func (bot *Bot) SendMarkdown(chatID, content string) (bool, error) {
	if len(strings.TrimSpace(chatID)) > 0 {
		bot.chatID = chatID
	}

	if len(strings.TrimSpace(bot.chatID)) < 1 {
		return false, errors.New("chatID cannot be empty")
	}
	urlTarget := prepareURLWith(bot.botToken, bot.chatID, content, "MarkdownV2")

	if bot.hTTPClient == nil {
		bot.hTTPClient = new(hTTPClient)
	}

	resp, err := bot.hTTPClient.Get(urlTarget, nil)
	defer resp.Body.Close()

	if err != nil {
		log.Printf("Request Error: [%v]\n", err)
		return false, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	log.Printf("Telegram response: %v\n", newStr)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Printf("Request Error. Response Body: [%s]\n", resp.Body)
		return false, fmt.Errorf("Request error: Status code [%d]", resp.StatusCode)
	}

	return true, nil
}

func prepareURLWith(token, chatID, content, parseMode string) string {
	baseURL, err := url.Parse(URLBase)
	if err != nil {
		log.Println("Malformed URL: ", err.Error())
	}

	baseURL.Path += "bot" + token + "/" + SendMessageEndpoint

	params := url.Values{}
	params.Add("chat_id", chatID)
	params.Add("text", content)
	params.Add("parse_mode", parseMode)

	baseURL.RawQuery = params.Encode() // Escape Query Parameters

	return baseURL.String()
}

func buildRootURLFrom(token string) string {
	return URLBase + "bot" + token
}
