package main

import (
	"go/format"
	"net/http"
	"os"
	"strings"

	"github.com/LlamaNite/llamalog"
	"github.com/haashemi/tgo/internal/tggen"
	"github.com/haashemi/tgo/internal/tgparser"
)

const TelegramDocURL = "https://core.telegram.org/bots/api"

var log = llamalog.NewLogger("TGO", "Generator")

var typesToIgnore = []string{"InputFile"}

func main() {
	resp, err := http.Get(TelegramDocURL)
	if err != nil {
		log.Fatal("Failed to fetch telegram doc > %s", err.Error())
	}
	defer resp.Body.Close()
	log.Info("Documentation fetched")

	groups, err := tgparser.Parse(resp.Body)
	if err != nil {
		log.Fatal("Failed to parse the data >> %s", err.Error())
	}
	log.Info("Groups parsed")

	var methods []string
	var types []string

	for _, g := range groups {
		if g.IsMethod() {
			methods = append(methods, tggen.GenerateMethod(g)...)
			log.Info("Generated Method %s", g.Name)
			continue
		}

		if shouldBeIgnored(g.Name, typesToIgnore) {
			continue
		}
		types = append(types, tggen.GenerateType(g))
		log.Info("Generated type %s", g.Name)
	}

	if err = formatAndSave("api_methods.go", methods); err != nil {
		log.Fatal("Failed to format and save methods >> %s", err.Error())
	}
	log.Info("Methods saved")

	if err = formatAndSave("api_types.go", types); err != nil {
		log.Fatal("Failed to format and save types >> %s", err.Error())
	}
	log.Info("Types saved")
}

func formatAndSave(path string, data []string) error {
	formattedCode, err := format.Source([]byte("package botapi\n\n" + strings.Join(data, "\n\n")))
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
