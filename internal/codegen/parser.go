package main

import (
	"fmt"
	"strings"
)

type StructField struct{ Name, LoweredName, Type, JsonTag, Description string }

func parseStructField(field Field) StructField {
	tag := field.Name
	if field.IsOptional {
		tag += ",omitempty"
	}

	desc := strings.ReplaceAll(field.Description, "\n", " ")
	if field.IsOptional && !strings.HasPrefix(desc, "Optional.") {
		desc = "Optional. " + desc
	}

	return StructField{
		Name:        snakeToPascal(field.Name),
		LoweredName: lowerFirstLetter(snakeToPascal(field.Name)),
		Type:        typeOf(field.Name, field.Type),
		JsonTag:     fmt.Sprintf("`json:\"%s\"`", tag),
		Description: desc,
	}
}

type Struct struct {
	Name, Description string
	EmbedTypes        []string
	Fields            []StructField
}

func parseStruct(g Group) (data Struct) {
	data.Name = g.Name
	data.Description = g.Description

	for _, field := range g.Fields {
		data.Fields = append(data.Fields, parseStructField(field))
	}

	return data
}

type Method struct {
	Name, RawName, Description, Params, ReturnType string
	EssentialParams, OptionalParams                []StructField
	UploadableParamsCheckCode                      string
}

func parseMethod(g Group) (data Method) {
	data.Name = upperFirstLetter(g.Name)
	data.RawName = g.Name
	data.Description = g.Description
	data.ReturnType = "json.RawMessage"
	if types := extractReturnType(g.Description); len(types) == 1 {
		data.ReturnType = typeOf(g.Name, types[0])
	}

	var uploadableParams []string
	for _, field := range g.Fields {
		structData := parseStructField(field)

		if field.IsOptional {
			data.OptionalParams = append(data.OptionalParams, structData)
		} else {
			data.EssentialParams = append(data.EssentialParams, structData)
		}

		if structData.Type == "InputFile" {
			uploadableParams = append(uploadableParams, "params."+structData.Name+".NeedsUpload()")
		}
	}

	var params []string
	for _, param := range data.EssentialParams {
		params = append(params, param.LoweredName+" "+param.Type)
	}

	if data.OptionalParams != nil {
		params = append(params, "optionalParams *"+data.Name+"Options")
	}

	data.Params = strings.Join(params, ", ")
	data.UploadableParamsCheckCode = strings.Join(uploadableParams, " || ")

	return data
}
