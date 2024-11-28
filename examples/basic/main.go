package main

import (
	"fmt"
	"log"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/filters"
	"github.com/haashemi/tgo/routers/message"
	"github.com/haashemi/tgo/tg"
)

const BotToken = "bot_token"

func main() {
	bot := tgo.NewBot(BotToken, tgo.Options{
		// it will set this parse mode for all api call via bot.Send, ctx.Send, and ctx.Reply
		DefaultParseMode: tg.ParseModeHTML,
	})

	// initialize a new router to handle messages
	messageRouter := message.NewRouter()
	// It will call the handler when a new message gets received with text/caption of "hi"
	messageRouter.Handle(filters.And(filters.IsMessage(), filters.Text("hi")), Hi)

	// Handlers are called in order (at the least in DefaultRouter, other routers may work differently)
	// so, if no handlers gets used and the update is a new message, Echo will be called.
	messageRouter.Handle(filters.And(filters.IsMessage()), Echo)

	// add our message router to the bot routers; so it will be triggered on updates.
	bot.AddRouter(messageRouter)

	botInfo, err := bot.GetMe()
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
func Hi(ctx *message.Context) {
	// Get sender's first name with getting the raw message
	senderFirstName := ctx.Message.From.FirstName

	// create the text using HTML Markups
	text := fmt.Sprintf("Hi <i>%s</i>!", senderFirstName)

	// HTML Parse mode will be automatically set
	ctx.Reply(&tg.SendMessage{
		Text: text,
	})
}

// Echo just echoes with text
func Echo(ctx *message.Context) {
	// get text or caption of the sent message and send it back!
	ctx.Send(&tg.SendMessage{Text: ctx.String()})
}
