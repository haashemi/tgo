package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"time"
)

const TelegramDocURL = "https://core.telegram.org/bots/api"

func main() {
	doc, err := Fetch()
	if err != nil {
		log.Fatalln("Failed to fetch the documentation >", err)
		return
	}

	startTime := time.Now()
	buf := bytes.NewBuffer(nil)
	err = Template.ExecuteTemplate(buf, "template.gtpl", Parse(doc))
	if err != nil {
		log.Fatalln("Failed to generate >", err)
		return
	}

	mewB, err := format.Source(buf.Bytes())
	if err != nil {
		os.WriteFile("api.gen.go", buf.Bytes(), os.ModePerm)
		return
	}
	os.WriteFile("api.gen.go", mewB, os.ModePerm)

	// bb, _ := json.MarshalIndent(Parse(doc), "", "    ")
	// os.WriteFile("data.json", bb, os.ModePerm)
	log.Println("generated in", time.Since(startTime))
}
