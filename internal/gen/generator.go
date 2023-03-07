package main

import (
	"io"
	"strings"
	"text/template"
)

var InterfaceGen = template.Must(template.New("interface-generator").Parse(`
{{range $desc := .Description }}
// {{ $desc }} {{end}}
type {{ .PascalCaseName }} interface {
	// Is{{ .PascalCaseName }} does nothing and is only used to enforce type-safety
	Is{{ .PascalCaseName }}()
}
`))

var TypeGen = template.Must(template.New("type-generator").Parse(`
{{range $desc := .Description }}
// {{ $desc }} {{end}}
type {{ .PascalCaseName }} struct {
	{{range $field := .EssentialFields -}}
		{{ $field.PascalCaseName }} {{ $field.Type }} {{ $field.Tag }} // {{ $field.Description }} 
	{{end -}}
	{{range $field := .OptionalFields -}}
		{{ $field.PascalCaseName }} {{ $field.Type }} {{ $field.Tag }} // {{ $field.Description }} 
	{{end -}}
}

{{ if .InterfaceName }}
func ({{ .PascalCaseName }}) Is{{ .InterfaceName }}() {}
{{ end }}
`))

var MethodGen = template.Must(template.New("method-generator").Parse(`
{{ if .OptionalFields }}
type {{ .PascalCaseName }}Options struct {
	{{range $field := .OptionalFields -}}
		{{ $field.PascalCaseName }} {{ $field.Type }} {{ $field.Tag }} // {{ $field.Description }} 
	{{end -}}	
}
{{end}}

type {{ .Name }}Params struct {
	{{ if .OptionalFields -}} *{{ .PascalCaseName }}Options {{ end }}

	{{range $field := .EssentialFields -}}
		{{ $field.PascalCaseName }} {{ $field.Type }} {{ $field.Tag }} // {{ $field.Description }} 
	{{end -}}
}

{{range $desc := .Description }}
// {{ $desc }} {{end}}
func (b *Client) {{ .PascalCaseName }}( {{ range $param := .EssentialFields -}} {{ $param.CamelCaseName }} {{ $param.Type }}, {{end -}} {{if .OptionalFields -}}optionalParams *{{ .PascalCaseName }}Options{{end -}}) ({{ .ReturnType }}, error) {
	params := &{{ .Name }}Params{}

	{{ range $param := .EssentialFields -}}
	params.{{ $param.PascalCaseName }} = {{ $param.CamelCaseName }}
	{{ end -}}

	{{ if .OptionalFields -}}
	params.{{ .PascalCaseName }}Options = optionalParams
	{{ end }}

	return doHTTP[{{ .ReturnType }}](b.client, b.url, "{{ .Name }}", params)
}

{{ if .Uploadables }}
func (params *{{ .Name }}Params) HasUploadable() bool {
	{{ range $up := .Uploadables }}
		if _, ok := params.{{ $up }}.(*InputFileUploadable); ok {
			return true
		}
	{{end}}
	return false
}
{{ end }}
`))

func Generate(data RawData, writer io.Writer) {
	// method validation
	if strings.ToLower(data.Name)[0] == data.Name[0] {
		MethodGen.Execute(writer, data)
		return
	}

	// Interface validation
	if !strings.Contains(data.Description[0], "holds no information") && data.OptionalFields == nil && data.EssentialFields == nil {
		InterfaceGen.Execute(writer, data)
		return
	}

	TypeGen.Execute(writer, data)
}
