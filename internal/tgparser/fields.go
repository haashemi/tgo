package tgparser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Field struct {
	Name        string
	Type        string
	Description string
	IsOptional  bool
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
