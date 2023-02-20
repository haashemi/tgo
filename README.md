# tgo

Golang Telegram Bot framework

this library is in early development and has a long way to its first stable release

## Usage

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/filter"
)

func main() {
	bot, err := tgo.NewBot("my-bot-token", tgo.Options{})
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle(filter.Text("hi"), HiHandler)

	fmt.Println("Bot is running as", bot.Me().Username)

	log.Fatal(bot.StartPolling())
}

func HiHandler(ctx *tgo.Context) {
	from := ctx.RawUpdate().Message.From
	fullName := from.FirstName + " " + from.LastName

	text := fmt.Sprintf("Hi <i>%s</i>\n\nHow old are you?", fullName)

	_, answer, err := ctx.Ask(
		tgo.NewText(text, tgo.ParseModeHTML).Options(tgo.TextOptions{
			ReplyMarkup: tgo.NewForceReply(),
		}),
		time.Second*10,
	)
	if err != nil {
		ctx.Send(tgo.NewText("you were too late... bye!"))
		return
	}

	answer.Reply(tgo.NewText("Wow! you are " + answer.Caption() + " years old?!"))
}

```
