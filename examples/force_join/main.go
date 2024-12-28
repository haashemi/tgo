package main

import (
	"flag"
	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/filters"
	"github.com/haashemi/tgo/tg"
	"log"
)

// ID of the channel that we want force join bot users to ir
const channelID int64 = -1002145897099

// bot's token must be passed by ./bot -token "my-awesome-token-here"
var botToken = flag.String("token", "", "Your Telegram bot's token")

// parsing the flags, nothing special
func init() { flag.Parse() }

func main() {
	bot := tgo.NewBot(*botToken, tgo.Options{})

	info, err := bot.API().GetMe()
	if err != nil {
		log.Fatalln("Failed to fetch the bot info", err)
	}

	bot.Handle(filters.Command("start", info.Username), startHandler, forceJoinMiddleware)

	log.Printf("Polling started as @%s\n", info.Username)
	log.Fatalln(bot.StartPolling(30))
}

// Disclaimer: it's not recommended to do GetChatMember on every single message you'll
// receive, as you may hit the Telegram's rate limit. which is not good...
//
// we recommend you to implement a basic caching or something.
func forceJoinMiddleware(next tgo.Handler) tgo.Handler {
	return func(ctx tgo.Context) {
		status, err := ctx.API().GetChatMember(&tg.GetChatMember{
			ChatId: tg.ID(channelID),
			UserId: ctx.From().Id,
		})
		if err != nil {
			_, _ = ctx.Send(&tgo.SendMessage{Text: "Failed to fetch your join status. try again later."})
			return
		}

		switch status.(type) {
		case *tg.ChatMemberOwner, *tg.ChatMemberAdministrator, *tg.ChatMemberMember:
			// they're joined! so everything's fine and we can call the next handler.
			next(ctx)
			return
		}

		// it has to be one of these: *tgo.ChatMemberLeft, *tgo.ChatMemberRestricted, or *tgo.ChatMemberBanned.
		_, _ = ctx.Send(&tgo.SendMessage{Text: "First you need to join to our channel somehow. try again after you joined"})
	}
}

func startHandler(ctx tgo.Context) {
	_, _ = ctx.Send(&tgo.SendMessage{Text: "Ok... Now you know how to implement a force-join thingy using tgo!"})
}
