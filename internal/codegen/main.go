package main

import (
	"go/format"
	"log"
	"net/http"
	"os"
	"strings"
)

const TelegramDocURL = "https://core.telegram.org/bots/api"

var typesToIgnore = []string{"InputFile"}

func main() {
	resp, err := http.Get(TelegramDocURL)
	if err != nil {
		log.Fatalf("Failed to fetch telegram doc > %s", err.Error())
	}
	defer resp.Body.Close()

	groups, err := ParseGroups(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse the data >> %s", err.Error())
	}

	var methods []string
	var types []string

	for _, g := range groups {
		if g.IsMethod() {
			methods = append(methods, GenerateMethod(g))
			continue
		}

		if shouldBeIgnored(g.Name, typesToIgnore) {
			continue
		}
		types = append(types, GenerateType(g))
	}

	if err = formatAndSave("api_methods.go", methods); err != nil {
		log.Fatalf("Failed to format and save methods >> %s", err.Error())
	}
	log.Println("Methods saved")

	if err = formatAndSave("api_types.go", types); err != nil {
		log.Fatalf("Failed to format and save types >> %s", err.Error())
	}
	log.Println("Types saved")
}

func formatAndSave(path string, data []string) error {
	formattedCode, err := format.Source([]byte("package tgo\n\n" + strings.Join(data, "\n\n")))
	if err != nil {
		return err
	}

	return os.WriteFile(path, formattedCode, os.ModePerm)
}

func shouldBeIgnored(s string, f []string) bool {
	for _, x := range f {
		if x == s {
			return true
		}
	}
	return false
}
