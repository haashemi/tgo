# TGO [Work in Progress]

TGo (hopefully) is an easy-to-use telegram bot framework.

It supports all [bot-api]("https://core.telegram.org/bots/api")'s methods and types with it's code-generator at `/internal/codegen` in `api_methods.go` & `api_types.go` out of the box.

The rest of this is package is a wrapper around the generated code to improve the developer experience and make it way easier to use.

## Basic example

```go
package main

import (
    "fmt"

    "github.com/haashemi/tgo"
    "github.com/haashemi/tgo/filters"
)

const ImageURL = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTK2nG24AYDm6FOEC7jIfgubO96GbRso2Xshu1f8abSYQ&s"

func main() {
    bot, _ := tgo.NewBot("bot:token", tgo.Options{})
    
    bot.OnMessage(filters.And(filters.IsMessage(), filters.Text("hi")), func(ctx tgo.Context) {
        ctx.Reply("hi!", nil)

        ctx.SendPhoto(tgo.FileFromURL(ImageURL), &tgo.SendPhotoOptions{
            Caption: "A random cute image",

            ReplyMarkup: tgo.NewReplyKeyboardMarkup(
                tgo.NewReplyKeyboardRow(tgo.NewKeyboardButton("hi")),
            ).
                WithResizeKeyboard().
                WithInputFieldPlaceholder("Say hi again..."),
        })
    })

    fmt.Println("Bot is running as", bot.Username)
    bot.StartPolling()
}
```
