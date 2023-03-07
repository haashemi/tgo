package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
)

const TelegramDocURL = "https://core.telegram.org/bots/api"

func main() {
	resp, err := FetchAndParse()
	if err != nil {
		log.Fatalln("failed to fetch or parse the data >>", err)
	}

	buf := bytes.NewBuffer(nil)
	buf.Write([]byte("// CODE GENERATED. DO NOT EDIT.\npackage tgo\n\n"))

	for _, data := range resp.Data {
		Generate(data, buf)
	}

	mewB, err := format.Source(buf.Bytes())
	if err != nil {
		os.WriteFile("api.go", buf.Bytes(), os.ModePerm)
		return
	} else {
		os.WriteFile("api.go", mewB, os.ModePerm)
	}
}
