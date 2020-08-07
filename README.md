## Telegrandma

[![Build Status](https://github.com/henriquefonseca/telegrandma/workflows/Go/badge.svg)](https://github.com/henriquefonseca/telegrandma/actions)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](go.mod)


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
    "github.com/henriquefonseca/telegrandma"
)

var (
    botToken = "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"
    chatID = "987654321"
)

func main() {
    bot, err := telegrandma.NewBot(botToken)
    if err != nil {
	    t.Errorf("Unexpected error: [%s]\n", err)
    }

    // Accessing bot last updates
    updates, err := bot.GetUpdates()
	if err != nil {
		t.Errorf("Unexpected error: [%s]\n", err)
	}

    log.Println("updates.Result", updates.Result)
    
    // Send a text message
    msg := "What happens at Nana's… stays at Nana's."
	if success, err := bot.SendMessage(ChatID, msg); !success {
		t.Errorf("Unexpected error: The message was not sent. Error: [%s]", err)
    }
    
    // Send a message with allowed html tags
    msg := "<b>If nothing is going well, call your grandmother.</b>"
	if success, err := bot.SendHTML(ChatID, msg); !success {
		t.Errorf("Unexpected error: The message was not sent. Error: [%s]", err)
	}

}
```

If you want to send html, see [html tags](https://core.telegram.org/bots/api#html-style) for more information about tags you can use.

## Contribution
Contributing is more than welcome. If you find any problem, please create an issue or send a pull request.


## License
Copyright (c) 2020-present [Henrique Fonseca](https://github.com/henriquefonseca)

Licensed under [MIT License](LICENSE)