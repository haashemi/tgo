package main

import (
	"fmt"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type TemplateData struct {
	Sections     []Section
	Implementers map[string]string
}

var Template = template.Must(template.New("tgo").Funcs(template.FuncMap{
	"append":               func(x []string, y string) []string { return append(x, y) },
	"snakeToPascal":        snakeToPascal,
	"upperFirstLetter":     upperFirstLetter,
	"lowerFirstLetter":     lowerFirstLetter,
	"extractReturnType":    extractReturnType,
	"getTag":               getTag,
	"getStringerMethod":    getStringerMethod,
	"getType":              getType,
	"getMediaFields":       getMediaFields,
	"getDefaultValue":      getDefaultValue,
	"isArray":              func(s string) bool { return strings.HasPrefix(s, "[]") },
	"getNestedMediaFields": getNestedMediaFields,
}).ParseFiles("./cmd/template.gtpl"))

func snakeToPascal(s string) string {
	return strings.ReplaceAll(cases.Title(language.English).String(strings.ReplaceAll(s, "_", "  ")), " ", "")
}

func upperFirstLetter(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func lowerFirstLetter(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}

func getTag(name string, isOptional bool) string {
	var optionalField string
	if isOptional {
		optionalField = ",omitempty"
	}

	return fmt.Sprintf("`json:\"%s%s\"`", name, optionalField)
}

func getMediaFields(s Section, sections []Section) (fields []Field) {
	for _, f := range s.Fields {
		if strings.HasSuffix(getType(f.Name, f.Type, f.IsOptional, sections), "InputFile") {
			fields = append(fields, f)
		}
	}

	return fields
}

func getNestedMediaFields(all []Section, s Section) (fields []Field) {
	for _, f := range s.Fields {
		ft := getType(f.Name, f.Type, f.IsOptional, all)
		ft = strings.TrimPrefix(ft, "[]")
		ft = strings.TrimPrefix(ft, "*")

		if ft == "InputMedia" {
			fields = append(fields, f)
			continue
		}

		for _, xs := range all {
			if xs.Name == ft {
				for _, xf := range xs.Fields {
					if getType(xf.Name, xf.Type, xf.IsOptional, all) == "*InputFile" {
						fields = append(fields, f)
						break
					}
				}
				break
			}
		}
	}

	return fields
}
