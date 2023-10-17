package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Section struct {
	Name        string
	Description []string
	Fields      []Field
	IsInterface bool
}

type Field struct {
	Name        string
	Type        string
	Description string
	IsOptional  bool
}

func Fetch() (*goquery.Document, error) {
	resp, err := http.Get(TelegramDocURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch telegram doc > %s", err.Error())
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func Parse(doc *goquery.Document) (data TemplateData) {
	data.Implementers = make(map[string]string)

	// looping through each group
	doc.Find("#dev_page_content h4").Each(func(i int, s *goquery.Selection) {
		// it doesn't seems to be a group, so we'll ignore it
		if strings.Contains(s.Text(), " ") {
			return
		}

		var section Section
		section.Name = s.Text()

		// We implement this ourselves.
		if section.Name == "InputFile" {
			return
		}

		s.NextUntil("h3,h4").Each(func(i int, s *goquery.Selection) {
			switch {
			case s.Is("p"), s.Is("blockquote"):
				section.Description = append(section.Description, strings.Split(s.Text(), "\n")...)

			case s.Is("ul"):
				supportedTypes := s.Find("li").Map(func(i int, s *goquery.Selection) string { return s.Text() })
				section.Description = append(section.Description, strings.Join(supportedTypes, ", "))

				for _, t := range supportedTypes {
					data.Implementers[t] = section.Name
				}
				if supportedTypes != nil {
					section.IsInterface = true
				}
			case s.Is("table"):
				s.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
					field := Field{}
					fields := s.Find("td").Map(func(i int, s *goquery.Selection) string { return s.Text() })

					field.Name = fields[0]
					field.Type = fields[1]

					switch len(fields) {
					case 3:
						field.IsOptional = strings.HasPrefix(fields[2], "Optional.")
						field.Description = strings.ReplaceAll(fields[2], "\n", " ")
					case 4:
						field.IsOptional = fields[2] == "Optional"
						field.Description = strings.ReplaceAll(fields[3], "\n", " ")
					}

					section.Fields = append(section.Fields, field)
				})
			}
		})

		// Normalize the description
		section.Description[0] = strings.Replace(section.Description[0], "This object", section.Name, 1)
		section.Description[0] = strings.Replace(section.Description[0], "Use this method to", section.Name+" is used to", 1)

		data.Sections = append(data.Sections, section)
	})

	// extended sections
	data.Sections = append(data.Sections, Section{
		Name:        "ReplyMarkup",
		Description: []string{"ReplyMarkup is an interface for InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, ForceReply"},
		IsInterface: true,
	})

	data.Sections = append(data.Sections, Section{
		Name:        "ChatID",
		Description: []string{"ChatID is an interface for usernames and chatIDs"},
		IsInterface: true,
	})

	// extended implementers
	data.Implementers["InlineKeyboardMarkup"] = "ReplyMarkup"
	data.Implementers["ReplyKeyboardMarkup"] = "ReplyMarkup"
	data.Implementers["ReplyKeyboardRemove"] = "ReplyMarkup"
	data.Implementers["ForceReply"] = "ReplyMarkup"

	return data
}
