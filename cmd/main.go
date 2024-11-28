package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"text/template"
	"time"
)

const TelegramDocURL = "https://core.telegram.org/bots/api"

type TemplateData struct {
	Sections     []Section
	Implementers map[string]string
}

var Template = template.Must(template.New("tgo").
	Funcs(template.FuncMap{
		"getInterfaceTemplate": getInterfaceTemplate,
		"getTypeTemplate":      getTypeTemplate,
		"getMethodTemplate":    getMethodTemplate,
		"isMethod":             isMethod,
		"isArray":              isArray,
	}).
	ParseFiles("./cmd/template.gotmpl"),
)

func main() {
	doc, err := Fetch()
	if err != nil {
		log.Fatalln("Failed to fetch the documentation >", err)
		return
	}

	startTime := time.Now()

	buf := bytes.NewBuffer(nil)
	parsedDoc := Parse(doc)

	err = Template.ExecuteTemplate(buf, "template.gotmpl", parsedDoc)
	if err != nil {
		log.Fatalln("Failed to generate >", err)
		return
	}

	if mewB, err := format.Source(buf.Bytes()); err != nil {
		os.WriteFile("tg/api_gen.go", buf.Bytes(), os.ModePerm)
	} else {
		os.WriteFile("tg/api_gen.go", mewB, os.ModePerm)
	}

	log.Println("generated in", time.Since(startTime))
}
