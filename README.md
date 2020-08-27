## Telegrandma

[![Build Status](https://github.com/henriquefonseca/telegrandma/workflows/Go/badge.svg)](https://github.com/henriquefonseca/telegrandma/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/henriquefonseca/telegrandma)](https://goreportcard.com/report/github.com/henriquefonseca/telegrandma)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/henriquefonseca/telegrandma)](https://pkg.go.dev/github.com/henriquefonseca/telegrandma)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](go.mod)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE)

A tiny wrapper around [Telegram](https://core.telegram.org/bots/api), that helps you use your telegram bots easily. So simple that even your grandmother could use it.

## Installation

```bash
go get -u github.com/henriquefonseca/telegrandma
```

## Code Example

After [creating the Telegram bot](https://core.telegram.org/bots#3-how-do-i-create-a-bot) you will receive a bot [token](https://core.telegram.org/bots/api#authorizing-your-bot).

For sending messages you will gonna need a [chatID](https://core.telegram.org/bots/api#chat).

So now you're ready to use Telegrandma:

```go
package main

import (
    "log"
    "github.com/henriquefonseca/telegrandma"
)

var (
    botToken = "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"
    chatID = "987654321"
)

func main() {
    bot, err := telegrandma.NewBot(botToken)
    if err != nil {
	    log.Printf("Unexpected error: [%s]\n", err)
    }

    // Accessing bot last updates
    updates, err := bot.GetUpdates()
	if err != nil {
		log.Printf("Unexpected error: [%s]\n", err)
	}

    log.Println("updates.Result", updates.Result)
    
    // Send a text message
    msg := "What happens at Nana's… stays at Nana's."
	if success, err := bot.SendMessage(chatID, msg); !success {
		log.Printf("Unexpected error: The message was not sent. Error: [%s]", err)
    }
    
    // Send a message with allowed html tags
    msg := "<b>If nothing is going well, call your grandmother.</b>"
	if success, err := bot.SendHTML(chatID, msg); !success {
		log.Printf("Unexpected error: The message was not sent. Error: [%s]", err)
    }
    
    // Send a message with allowed markdown tags
    `Grandma is cooking 😋

	🎂 cake
	🍞 bread
	🍝 spaghetti
	🍦 ice cream

	* So hungry *`
	if success, err := bot.SendMarkdown(chatID, msg); !success {
		log.Printf("Unexpected error: The message was not sent. Error: [%s]", err)
	}

}
```

If you want to send html, see [html tags](https://core.telegram.org/bots/api#html-style) for more information about tags you can use.

If you want to send markdown, see [markdown tags](https://core.telegram.org/bots/api#markdown-style) for more information about tags you can use.

## Contribution
Contributing is more than welcome. If you find any problem, please create an issue or send a pull request.


## License
Copyright (c) 2020-present [Henrique Fonseca](https://github.com/henriquefonseca)

Licensed under [MIT License](LICENSE)