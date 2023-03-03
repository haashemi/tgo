package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"os"
)

const TelegramDocURL = "https://core.telegram.org/bots/api"

func main() {
	x, _ := FetchAndParse()
	b, _ := json.MarshalIndent(x, "", "    ")
	os.WriteFile("schema.json", b, os.ModePerm)

	buf := bytes.NewBuffer(nil)
	buf.Write([]byte("// CODE GENERATED. DO NOT EDIT.\npackage tgo\n\n"))

	for _, data := range x.Data {
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
