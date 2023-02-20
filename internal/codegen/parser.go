package main

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Field struct {
	Name        string
	Type        string
	Description string
	IsOptional  bool
}

type Group struct {
	Name        string
	Description string
	Fields      []Field
}

func ParseGroups(data io.Reader) ([]Group, error) {
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return nil, err
	}

	return extractGroups(doc.Find("#dev_page_content h4")), nil
}

func (g Group) IsMethod() bool {
	return strings.ToLower(g.Name)[0] == g.Name[0]
}

func extractGroups(s *goquery.Selection) (groups []Group) {
	s.Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), " ") {
			return
		}

		group := Group{
			Name:        s.Text(),
			Description: strings.ReplaceAll(s.NextFiltered("p").Text(), "\n", " "),
		}
		extractFieldsTo(&group, s)

		groups = append(groups, group)
	})

	return groups
}

func extractFieldsTo(g *Group, s *goquery.Selection) {
	s.NextUntil("h4").Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		field := s.Find("td").Map(func(i int, s *goquery.Selection) string { return s.Text() })

		if g.IsMethod() {
			g.Fields = append(g.Fields, Field{
				Name:        field[0],
				Type:        field[1],
				Description: field[3],
				IsOptional:  field[2] == "Optional",
			})
		} else {
			g.Fields = append(g.Fields, Field{
				Name:        field[0],
				Type:        field[1],
				Description: field[2],
				IsOptional:  strings.HasPrefix(field[2], "Optional."),
			})
		}
	})
}
