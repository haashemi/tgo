package main

import "strings"

type InterfaceTemplate struct {
	Name        string
	Description string
	InterfaceOf []string
}

func getInterfaceTemplate(section Section, sections []Section) InterfaceTemplate {
	return InterfaceTemplate{
		Name:        section.Name,
		Description: "// " + strings.Join(section.Description, "\n// "),
		InterfaceOf: section.InterfaceOf,
	}
}
