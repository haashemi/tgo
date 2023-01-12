package tgparser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Group struct {
	Name        string
	Description string
	Fields      []Field
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
