package main

import (
	"fmt"
	"strings"
)

type MethodTemplate struct {
	Name                  string
	MethodName            string
	ReturnType            string
	ReturnsInterface      bool
	Description           string
	EmptyReturnValue      string
	Fields                []MethodField
	InputFileFields       []MethodField
	NestedInputFileFields []MethodField
}

type MethodField struct {
	Name          string
	FieldName     string
	Type          string
	Tag           string
	Description   string
	DefaultValue  string
	StringifyCode string
	IsOptional    bool
}

func NewMethodFields(fields []Field, sections []Section) (newFields []MethodField, inputFileFields []MethodField) {
	for _, f := range fields {
		var optionalField string
		if f.IsOptional {
			optionalField = ",omitempty"
		}

		tf := MethodField{
			Name:        snakeToPascal(f.Name),
			FieldName:   f.Name,
			Type:        getType(f.Name, f.Type, f.IsOptional, sections),
			Tag:         fmt.Sprintf("`json:\"%s%s\"`", f.Name, optionalField),
			Description: f.Description,
			IsOptional:  f.IsOptional,
		}
		tf.DefaultValue = defaultValueOfType(tf.Type)
		tf.StringifyCode = stringifyTypeAndSetInMap(tf.Name, tf.FieldName, tf.Type)

		if tf.Type == "*InputFile" {
			inputFileFields = append(inputFileFields, tf)
		}

		newFields = append(newFields, tf)
	}

	return
}

func getMethodTemplate(section Section, sections []Section) MethodTemplate {
	fields, inputFileFields := NewMethodFields(section.Fields, sections)
	nestedInputFileFields, _ := NewMethodFields(getNestedMediaFields(sections, section), sections)
	returnType := getType("", extractReturnType(section.Description), true, sections)

	desc := strings.Join(section.Description, "\n// ")
	if desc != "" {
		desc = "// " + desc
	}

	mt := MethodTemplate{
		Name:                  strings.ToUpper(string(section.Name[0])) + section.Name[1:],
		MethodName:            section.Name,
		ReturnType:            returnType,
		ReturnsInterface:      isInterface(returnType, sections),
		Description:           desc,
		EmptyReturnValue:      defaultValueOfType(returnType),
		Fields:                fields,
		InputFileFields:       inputFileFields,
		NestedInputFileFields: nestedInputFileFields, // TODO
	}

	return mt
}
