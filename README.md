# tgo

Golang Telegram Bot framework

this library is in early development and has a long way to its first stable release

## Usage

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
    
    bot.OnMessage(filters.Text("hi"), func(ctx tgo.MessageContext) {
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
