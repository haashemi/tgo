package main

import (
	"fmt"
	"log"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/filters"
)

const BotToken = "bot_token"

func main() {
	bot, err := tgo.NewBot(BotToken, tgo.Options{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// It will call the handler when a new message gets received with text/caption of "hi"
	bot.Handle(filters.And(filters.IsMessage(), filters.Text("hi")), Hi)

	// Handlers are called in order (at the least in DefaultRouter, other routers may work differently)
	// so, if no handlers gets used and the update is a new message, Echo will be called.
	bot.Handle(filters.And(filters.IsMessage()), Echo)

	log.Println("Bot is started as", bot.Username)

	if err := bot.StartPolling(); err != nil {
		log.Fatalln("Polling stopped >>", err.Error())
		return
	}
	log.Println("Bot stopped successfully")
}

// Hi answers the hi message with a new hi!
func Hi(ctx tgo.Context) {
	// Get sender's first name with getting the raw message
	senderFirstName := ctx.Message.From.FirstName

	// create the text using HTML Markups
	text := fmt.Sprintf("Hi <i>%s</i>!", senderFirstName)

	ctx.Reply(text, &tgo.SendMessageOptions{
		// pass the parse mode, so telegram knows that our text contains HTML Markup!
		ParseMode: tgo.ParseModeHTML,
	})
}

// Echo just echoes with text
func Echo(ctx tgo.Context) {
	// get text or caption of the sent message and send it back!
	ctx.Send(ctx.Text(), nil)
}
