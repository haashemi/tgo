# TGO

[![Go Reference](https://pkg.go.dev/badge/github.com/haashemi/tgo.svg)](https://pkg.go.dev/github.com/haashemi/tgo)

TGO is a simple, flexible, and fully featured telegram-bot-api framework for Go developers.

It gives you the ability to implement your own filters, middlewares, and even routers!

All API methods and types are all code generated from the [telegram's documentation](https://core.telegram.org/bots/api) in `./cmd` at the `api.gen.go` file.

## Installation

As tgo is currently work-in-progress, we recommend you be on the latest git commit instead of git release.

```bash
$ go get -u github.com/haashemi/tgo@main
```

## Usage

More detailed examples can be found in the [examples folder](/examples/)

> Help us add more useful examples by doing a PR!

### Basic

```go
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
	bot := tgo.NewBot(BotToken, tgo.Options{ DefaultParseMode: tgo.ParseModeHTML })

	info, err := bot.GetMe()
	if err != nil {
		log.Fatalln("Failed to fetch the bot info", err.Error())
	}

	mr := message.NewRouter()
	mr.Handle(filters.Command("start", info.Username), Start)
	bot.AddRouter(mr)

	for {
		log.Println("Polling started as", info.Username)

		if err := bot.StartPolling(30, "message"); err != nil {
			log.Fatalln("Polling failed >>", err.Error())
			time.Sleep(time.Second * 5)
		}
	}
}

func Start(ctx *message.Context) {
	text := fmt.Sprintf("Hi <i>%s</i>!", ctx.Message.From.FirstName)

    ctx.Reply(&tgo.SendMessage{Text: text})
}
```

### Uploading a file

```go
// FileID of the photo that you want to reuse.
photo := tgo.FileFromID("telegram-file-id")

// Image URL from somewhere else.
photo = tgo.FileFromURL("https://cataas.com/cat")

// Local file path, used when you ran telegram-bot-api locally.
photo = tgo.FileFromPath("/home/tgo/my-nsfw-image.png")

// When you want to upload and image by yourself,
// which is something that you'll usually do.
photo = tgo.FileFromReader("my-awesome-image.png", reader)

bot.Send(&tgo.SendPhoto{
	ChatId: tgo.ID(0000000),
	Photo: photo,
})
```

### Ask a Question

It happens when you want to ask a question from the user and wait a few seconds for their response.

```go
// the question we want to ask.
var msg tgo.Sendable = &tgo.SendMessage{ Text: "How old are you, cutie?" }

// Here, we'll ask the question from the user (userId) in the chat (chatId)
// and wait for their response for 30 seconds.
//
// if the user don't responds in 30 seconds, context.DeadlineExceeded error
// will be returned.
question, answer, err := bot.Ask(chatId, userId, msg, time.Seconds*30)
if err != nil {
    // handle the error
}
// do whatever you want with the Q & A.
```

### Sessions

You may want to store some temporarily in-memory data for the user, tgo's Bot Session got you.

```go
// simple as that! it will return a *sync.Map
session := bot.GetSession(userID)

// using message or callback router:
session := ctx.Session()
```

### Routers

As you've read from the beginning, you're able to implement your own routers in the way you want. There are currently three built-in routers which are `message`, `callback`, and the raw update handlers.

Read more in the [routers section](/routers/)

### Polling

Polling is the easiest part of the bot. (it has to be.)
You just add your routers using [bot.AddRouter](https://pkg.go.dev/github.com/haashemi/tgo#Bot.AddRouter), and just do [bot.StartPolling](https://pkg.go.dev/github.com/haashemi/tgo#Bot.StartPolling). Here's an example:

```go
// initialize the bot
bot := tgo.NewBot("...", tgo.Options{})

// initialize our routers
mr := messages.NewRouter()
mr.Handle(myFilter1, myHandler1)
mr.Handle(myFilter2, myHandler2)

mrPrivate := messages.NewRouter(mySecurityMiddleware)
mrPrivate.Handle(myFilter, myPrivateHandler)

cr := callback.NewRouter()
cr.Handle(myFilter3, myCallbackHandler)

// add our routers to the bot in the order we want
bot.AddRouter(mr)
bot.AddRouter(mrPrivate)
bot.AddRouter(cr)

// start polling with the timeout of 30 seconds
// and only listen for message and callback_query updates.
err := bot.StartPolling(30, "message", "callback_query")
// handle the errors here.
```

## Contributions

1. Open an issue and describe what you're gonna do.
2. Fork, Clone, and do a Pull Request!

We are in our beta stage and there are a lot of thing to do. With that being said, all type of contributions are highly appreciated whatever it would be adding new features, fixing bugs, writing tests or docs, improving the current docs' grammars, or even fixing the typos!
