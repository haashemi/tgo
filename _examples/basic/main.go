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

	bot.OnMessage(filters.Text("hi"), Echo)

	bot.OnMessage(filters.True(), Echo)

	log.Println("Bot is started as", bot.Username)
	log.Fatalln(bot.StartPolling())
}

// Hi answers the hi message with a new hi!
func Hi(ctx tgo.MessageContext) {
	// Get sender's first name with getting the raw message
	senderFirstName := ctx.Message().From.FirstName

	// create the text using HTML Markups
	text := fmt.Sprintf("Hi <i>%s</i>!", senderFirstName)

	ctx.Send(text, &tgo.SendMessageOptions{
		// pass the parse mode, so telegram knows that our text contains HTML Markup!
		ParseMode: tgo.ParseModeHTML,
	})
}

// Echo just echoes with text
func Echo(ctx tgo.MessageContext) {
	// get text or caption of the sent message
	text := ctx.Caption()

	// send it back!
	ctx.Send(text, nil)
}
