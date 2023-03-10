package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"time"
)

const TelegramDocURL = "https://core.telegram.org/bots/api"

func main() {
	s := time.Now()

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

	fmt.Println("generated in", time.Since(s))
}
