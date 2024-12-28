package main

import (
	"fmt"
	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/filters"
	"github.com/haashemi/tgo/tg"
	"log"
	"time"
)

const BotToken = "bot_token"

func main() {
	bot := tgo.NewBot(BotToken, tgo.Options{
		// it will set this parse mode for all api call via bot.Send, ctx.Send, and ctx.Reply
		DefaultParseMode: tg.ParseModeHTML,
	})

	info, err := bot.API().GetMe()
	if err != nil {
		log.Fatalln("Failed to fetch the bot info", err.Error())
	}

	// register a handler for /start command, which also works for groups.
	bot.Handle(filters.Command("start", info.Username), Start)

	// start polling in an infinite loop
	for {
		log.Println("Polling started as", info.Username)

		// start the long-polling with the timeout of 30 seconds
		// and only new messages are allowed as an update (to save traffic or whatever).
		if err := bot.StartPolling(30, "message"); err != nil {
			log.Println("Polling failed >>", err.Error())
			log.Println("Sleeping for 5 seconds...")
			time.Sleep(time.Second * 5)
		}
	}
}

// Start says hi to the user!
func Start(ctx tgo.Context) {
	// Get the message from the context.
	// We're sure that it's a message, so we don't have to do any nil check.
	msg := ctx.Message()

	// Get sender's first name with getting the raw message
	senderFirstName := msg.From.FirstName

	// create the text using HTML Markups
	text := fmt.Sprintf("Hi <i>%s</i>!", senderFirstName)

	// HTML Parse mode will be automatically set
	_, _ = ctx.Reply(&tgo.SendMessage{Text: text})
}
