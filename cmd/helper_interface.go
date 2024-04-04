package main

import "strings"

type InterfaceTemplate struct {
	Name        string
	Description string
	InterfaceOf []string
}

func getInterfaceTemplate(section Section, sections []Section) InterfaceTemplate {
	desc := strings.Join(section.Description, "\n// ")
	if desc != "" {
		desc = "// " + desc
	}

	return InterfaceTemplate{
		Name:        section.Name,
		Description: desc,
		InterfaceOf: section.InterfaceOf,
	}
}
