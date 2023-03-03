package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type RawResponse struct {
	Data            []RawData
	InterfaceFields map[string]string
}

type RawData struct {
	// Raw type/method name
	Name string

	// Raw name, but first letter is upperCased
	PascalCaseName string

	// full description containing notes and quotes.
	Description []string

	// METHOD ONLY: return type of the method
	ReturnType string

	// METHOD ONLY:
	Uploadables []string

	// INTERFACE ONLY: supported types of an interface
	SupportedTypes []string

	// STRUCT ONLY: name of the interface that supports this type
	InterfaceName string

	EssentialFields []RawField
	OptionalFields  []RawField
}

type RawField struct {
	// Raw field name
	Name string

	// Raw name, but first letter is upperCased
	PascalCaseName string

	// Raw name, but first letter is lowerCased
	CamelCaseName string

	Tag string

	// Type is Golang normalized type
	Type string

	IsOptional bool

	Description string
}

func FetchAndParse() (*RawResponse, error) {
	resp, err := http.Get(TelegramDocURL)
	if err != nil {
		log.Fatalf("Failed to fetch telegram doc > %s", err.Error())
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	raw := &RawResponse{InterfaceFields: map[string]string{}}

	// all groups are located here.
	// and if there is any empty space between them, they are most linky not a group
	doc.Find("#dev_page_content h4").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), " ") {
			return
		}

		data := RawData{Name: s.Text(), PascalCaseName: upperFirstLetter(s.Text()), InterfaceName: raw.InterfaceFields[s.Text()]}

		s.NextUntil("h3,h4").Each(func(i int, s *goquery.Selection) {
			switch {
			case s.Is("p"), s.Is("blockquote"):
				data.Description = append(data.Description, strings.Split(s.Text(), "\n")...)

			case s.Is("ul"):
				data.SupportedTypes = s.Find("li").Map(func(i int, s *goquery.Selection) string { return s.Text() })
				data.Description = append(data.Description, strings.Join(data.SupportedTypes, ", "))

				for _, t := range data.SupportedTypes {
					raw.InterfaceFields[t] = data.Name
				}

			case s.Is("table"):
				s.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
					fields := s.Find("td").Map(func(i int, s *goquery.Selection) string { return s.Text() })

					field := RawField{
						Name:           fields[0],
						PascalCaseName: snakeToPascal(fields[0]),
					}
					field.CamelCaseName = lowerFirstLetter(field.PascalCaseName)

					switch len(fields) {
					case 3:
						field.IsOptional = strings.HasPrefix(fields[2], "Optional.")
						field.Description = strings.ReplaceAll(fields[2], "\n", " ")
					case 4:
						field.IsOptional = fields[2] == "Optional"
						field.Description = strings.ReplaceAll(fields[3], "\n", " ")
					}

					field.Tag = getStructFieldTag(field.Name, field.IsOptional)

					field.Type = typeOf(fields[0], fields[1], field.IsOptional)
					if field.Type == "InputFile" {
						data.Uploadables = append(data.Uploadables, field.PascalCaseName)
					}

					if field.IsOptional {
						data.OptionalFields = append(data.OptionalFields, field)
					} else {
						data.EssentialFields = append(data.EssentialFields, field)
					}
				})
			}
		})

		// Extract the return type from the first description when all of them extracted
		data.ReturnType = typeOf("", extractReturnType(data.Description[0]), true)

		// Make type and function's descriptions more readable
		data.Description[0] = strings.Replace(data.Description[0], "This object", data.Name, 1)
		data.Description[0] = strings.Replace(data.Description[0], "Use this method to", data.Name+" is used to", 1)

		raw.Data = append(raw.Data, data)
	})

	return raw, nil
}
