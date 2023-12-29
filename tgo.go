// TGO is a simple, flexible, and fully featured telegram-bot-api framework for Go developers.
//
// It gives you the ability to implement your own filters, middlewares, and even routers!
//
// All API methods and types are all code generated from the [telegram's documentation] in `./cmd` at the `api.gen.go` file.
//
// Example:
//
//	package main
//
//	import (
//		"fmt"
//		"log"
//		"time"
//
//		"github.com/haashemi/tgo"
//		"github.com/haashemi/tgo/filters"
//		"github.com/haashemi/tgo/routers/message"
//	)
//
//	const BotToken = "bot_token"
//
//	func main() {
//		bot := tgo.NewBot(BotToken, tgo.Options{ DefaultParseMode: tgo.ParseModeHTML })
//
//		info, err := bot.GetMe()
//		if err != nil {
//			log.Fatalln("Failed to fetch the bot info", err.Error())
//		}
//
//		mr := message.NewRouter()
//		mr.Handle(filters.Command("start", info.Username), Start)
//		bot.AddRouter(mr)
//
//		for {
//			log.Println("Polling started as", info.Username)
//
//			if err := bot.StartPolling(30, "message"); err != nil {
//				log.Fatalln("Polling failed >>", err.Error())
//				time.Sleep(time.Second * 5)
//			}
//		}
//	}
//
//	func Start(ctx *message.Context) {
//		text := fmt.Sprintf("Hi <i>%s</i>!", ctx.Message.From.FirstName)
//
//		ctx.Reply(&tgo.SendMessage{Text: text})
//	}
//
// [telegram's documentation]: https://core.telegram.org/bots/api
package tgo
