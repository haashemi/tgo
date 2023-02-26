package main

import (
	"fmt"
	"os"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/filters"
)

func main() {
	bot, _ := tgo.NewBot("5880190686:AAGVvRUPz0OAVDwhDgXXx5xzmFBFCjadQsk", tgo.Options{})

	RegisterAdminHandlers(bot.Party(filters.And(filters.Text("uwu"))))

	fmt.Println("Running as ", bot.Username)
	bot.StartPolling()
}

func RegisterAdminHandlers(party tgo.Party) {
	party.Handle(filters.Text("hi"), func(ctx tgo.Context) {
		ctx.Send("hi man", nil)
	})

	party.Handle(filters.True(), EchoHandler)
}

func StartHandler(ctx tgo.Context) {

	ctx.Bot().GetFile("file-id")

	ctx.Reply("Hi!\nI'm EchoBot!\nPowered by aiogram.", nil)
}

func CatHandler(ctx tgo.Context) {
	photo, _ := os.Open("data/cats.jpg")

	ctx.ReplyPhoto(
		tgo.FileFromReader("image-of-cats.jpg", photo),
		&tgo.SendPhotoOptions{
			Caption: "Cats are here 😺",
		},
	)
}

func EchoHandler(ctx tgo.Context) {
	ctx.Send(ctx.Text(), nil)
}
