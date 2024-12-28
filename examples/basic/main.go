package main

import (
	"fmt"
	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/filters"
	"github.com/haashemi/tgo/tg"
	"log"
)

const BotToken = "bot_token"

func main() {
	bot := tgo.NewBot(BotToken, tgo.Options{
		// it will set this parse mode for all api call via bot.Send, ctx.Send, and ctx.Reply
		DefaultParseMode: tg.ParseModeHTML,
	})

	// It will call the handler when a new message gets received with text/caption of "hi"
	bot.Handle(filters.And(filters.IsMessage(), filters.Text("hi")), Hi)

	// Handlers are called in order (at the least in DefaultRouter, other routers may work differently)
	// so, if no handlers gets used and the update is a new message, Echo will be called.
	bot.Handle(filters.And(filters.IsMessage()), Echo)

	botInfo, err := bot.API().GetMe()
	if err != nil {
		log.Fatalln("Failed to fetch the bot info", err.Error())
		return
	}
	log.Println("Bot is started as", botInfo.Username)

	// start the long-polling with the timeout of 30 seconds
	// and only new messages are allowed as an update (to save traffic or whatever).
	if err := bot.StartPolling(30, "message"); err != nil {
		log.Fatalln("Polling stopped >>", err.Error())
		return
	}
	log.Println("Bot stopped successfully")
}

// Hi answers the hi message with a new hi!
func Hi(ctx tgo.Context) {
	// Get sender's first name with getting the raw message
	senderFirstName := ctx.From().FirstName

	// create the text using HTML Markups
	text := fmt.Sprintf("Hi <i>%s</i>!", senderFirstName)

	// HTML Parse mode will be automatically set
	_, _ = ctx.Reply(&tgo.SendMessage{Text: text})
}

// Echo just echoes with text
func Echo(ctx tgo.Context) {
	msg := ctx.Message()

	text := msg.Text
	if text == "" {
		text = msg.Caption
	}

	// get text or caption of the sent message and send it back!
	if text != "" {
		_, _ = ctx.Send(&tgo.SendMessage{Text: text})
	}
}
