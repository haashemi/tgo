package tgparser

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

func Parse(data io.Reader) ([]Group, error) {
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return nil, err
	}

	return extractGroups(doc.Find("#dev_page_content h4")), nil
}
