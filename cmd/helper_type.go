package main

import (
	"fmt"
	"strings"
)

type TypeTemplate struct {
	Name              string
	Description       string
	Implements        string
	Fields            []TypeField
	ContainsInterface bool
	ContainsInputFile bool
}

type TypeField struct {
	Name        string
	FieldName   string
	Type        string
	Tag         string
	Description string
	IsInterface bool
	IsInputFile bool
}

func NewTypeFields(fields []Field, sections []Section) (newFields []TypeField, containsInterface, containsInputFile bool) {
	for _, f := range fields {
		var optionalField string
		if f.IsOptional {
			optionalField = ",omitempty"
		}

		tf := TypeField{
			Name:        snakeToPascal(f.Name),
			FieldName:   f.Name,
			Type:        getType(f.Name, f.Type, f.IsOptional, sections),
			Tag:         fmt.Sprintf("`json:\"%s%s\"`", f.Name, optionalField),
			Description: f.Description,
		}
		tf.IsInputFile = tf.Type == "*InputFile"

		if !tf.IsInputFile {
			for _, s := range sections {
				if s.Name == tf.Type && s.IsInterface {
					tf.IsInterface = true
					break
				}
			}
		}

		newFields = append(newFields, tf)
	}

	for _, f := range newFields {
		if f.IsInterface {
			containsInterface = true
			break
		}
	}

	for _, f := range newFields {
		if f.IsInputFile {
			containsInputFile = true
			break
		}
	}

	return
}

func getTypeTemplate(section Section, sections []Section, implementers map[string]string) TypeTemplate {
	fields, containsInterface, containsInputFile := NewTypeFields(section.Fields, sections)

	desc := strings.Join(section.Description, "\n// ")
	if desc != "" {
		desc = "// " + desc
	}

	return TypeTemplate{
		Name:              section.Name,
		Description:       desc,
		Implements:        implementers[section.Name],
		Fields:            fields,
		ContainsInterface: containsInterface,
		ContainsInputFile: containsInputFile,
	}
}
