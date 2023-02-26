package main

import (
	"bytes"
	"text/template"
)

var TypeGen = template.Must(template.New("type-generator").Parse(`
// {{ .Name }} {{ .Description }}
type {{ .Name }} struct {
	{{range $embedType := .EmbedTypes -}}
		{{ $embedType }}
	{{end -}}
	
	{{range $param := .Fields -}}
		{{ $param.Name }} {{ $param.Type }} {{ $param.JsonTag }} // {{ $param.Description }} 
	{{end -}}
}
`))

var MethodGen = template.Must(template.New("method-generator").Parse(`
// {{ .Name }} {{ .Description }}
func (b *Bot) {{ .Name }}( {{ .Params }} ) ({{ .ReturnType }}, error) {
	params := &{{ .RawName }}Params{}

	{{ range $param := .EssentialParams -}}
	params.{{ $param.Name }} = {{ $param.LoweredName }}
	{{ end -}}

	{{ if .OptionalParams -}}
	params.{{ .Name }}Options = optionalParams
	{{ end }}

	return doHTTP[{{ .ReturnType }}](b.client, b.url, "{{ .RawName }}", params)
}

{{ if .UploadableParamsCheckCode }}
func (params *{{ .RawName }}Params) HasUploadable() bool {
	return {{ .UploadableParamsCheckCode }}
}
{{ end }}
`))

func GenerateType(g Group) string {
	buf := bytes.NewBuffer(nil)
	TypeGen.Execute(buf, parseStruct(g))
	return buf.String()
}

func GenerateMethod(g Group) string {
	buf := bytes.NewBuffer(nil)
	data := parseMethod(g)

	var embedField string
	if data.OptionalParams != nil {
		TypeGen.Execute(buf, Struct{
			Name:        data.Name + "Options",
			Description: data.Name + "Options contains " + data.Name + "'s optional params",
			Fields:      data.OptionalParams,
		})
		embedField = "*" + data.Name + "Options"
	}

	TypeGen.Execute(buf, Struct{
		Name:        data.RawName + "Params",
		Description: data.RawName + "Params contains " + data.Name + "'s params",
		EmbedTypes:  []string{embedField},
		Fields:      data.EssentialParams,
	})

	MethodGen.Execute(buf, data)

	return buf.String()
}
