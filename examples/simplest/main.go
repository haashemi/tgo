package main

import (
	"fmt"
	"log"
	"time"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/filters"
	"github.com/haashemi/tgo/routers/message"
)

const BotToken = "bot_token"

func main() {
	bot := tgo.NewBot(BotToken, tgo.Options{
		// it will set this parse mode for all api call via bot.Send, ctx.Send, and ctx.Reply
		DefaultParseMode: tgo.ParseModeHTML,
	})

	info, err := bot.GetMe()
	if err != nil {
		log.Fatalln("Failed to fetch the bot info", err.Error())
	}

	// initialize a new router to handle messages
	mr := message.NewRouter()

	// register a handler for /start command, which also works for groups.
	mr.Handle(filters.Command("start", info.Username), Start)

	// add our message router to the bot routers; so it will be triggered on updates.
	bot.AddRouter(mr)

	// start polling in an infinite loop
	for {
		log.Println("Polling started as", info.Username)

		// start the long-polling with the timeout of 30 seconds
		// and only new messages are allowed as an update (to save traffic or whatever).
		if err := bot.StartPolling(30, "message"); err != nil {
			log.Fatalln("Polling failed >>", err.Error())
			log.Println("Sleeping for 5 seconds...")
			time.Sleep(time.Second * 5)
		}
	}
}

// Start says hi to the user!
func Start(ctx *message.Context) {
	// Get sender's first name with getting the raw message
	senderFirstName := ctx.Message.From.FirstName

	// create the text using HTML Markups
	text := fmt.Sprintf("Hi <i>%s</i>!", senderFirstName)

	// HTML Parse mode will be automatically set
	ctx.Reply(&tgo.SendMessage{Text: text})
}
