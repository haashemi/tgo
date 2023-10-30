{{ define "header" }}
	// CODE GENERATED. DO NOT EDIT.
	package tgo

	import (
		"fmt"
		"errors"
		"strconv"
		"encoding/json"
	)
{{ end }}

{{ define "interface" }}
	{{ .Description }}
	type {{ .Name }} interface {	
		// Is{{ .Name }} does nothing and is only used to enforce type-safety
		Is{{ .Name }}()

		{{ if eq .Name "InputMedia" }}getFiles() map[string]*InputFile{{ end }}
	}

	{{ if .InterfaceOf }}
		func unmarshal{{.Name}}(rawBytes json.RawMessage) (data {{.Name}}, err error) {
			{{ range $index, $type := .InterfaceOf -}}
				data{{$type}} := &{{$type}}{}
				if err = json.Unmarshal(rawBytes, data{{$type}}); err == nil {
					return data{{$type}}, nil
				}{{if eq $index 0 -}} 
					else if _, ok := err.(*json.UnmarshalTypeError); err != nil && !ok {
						return nil, err
					}
				{{end-}}
			{{ end }}

			return nil, errors.New("unknown type")
		}
	{{ end }}
{{ end }}

{{ define "type" }}
	{{ .Description }}
	type {{ .Name }} struct { {{range $field := .Fields -}}
		{{ $field.Name }} {{ $field.Type }} {{ $field.Tag }} // {{ $field.Description }} 
	{{end -}} }

	{{if .Implements }}
		func ({{ .Name }}) Is{{ .Implements }}() {}
	{{end}}

	{{ if .ContainsInterface }}
		func (x *{{ .Name }}) UnmarshalJson(rawBytes []byte) (err error) {
			if len(rawBytes) == 0 {
				return nil, nil
			}

			type temp struct {
				{{range $field := .Fields -}}
					{{ $field.Name }} {{ if $field.IsInterface }}json.RawMessage{{ else }}{{ $field.Type }}{{ end }} {{ $field.Tag }} // {{ $field.Description }} 
				{{end -}}
			}
			raw := &temp{}

			if err = json.Unmarshal(rawBytes, raw); err != nil {
				return err
			}

			{{range $field := .Fields -}}{{ if $field.IsInterface }}
			if data, err := unmarshal{{$field.Type}}(raw.{{$field.Name}}); err != nil {
				return err
			} else {
				x.{{$field.Name}} = data
			}
			{{end }}{{end -}}
			{{range $field := .Fields -}}
				{{ if not $field.IsInterface }}x.{{ $field.Name }} = raw.{{ $field.Name }}{{ end }}
			{{end -}}

			return nil
		}
	{{ end }}

	{{ if .ContainsInputFile }}
		func (x *{{ .Name }}) getFiles() map[string]*InputFile {
			media := map[string]*InputFile{}

			{{ range $field := .Fields -}}
				{{ if $field.IsInputFile -}}
					if x.{{ $field.Name }} != nil {
						if x.{{ $field.Name }}.IsUploadable() {
							media["{{ $field.FieldName }}"] = x.{{ $field.Name }}
						}
					}
				{{ end -}}
			{{ end }}

			return media
		}
	{{ end }}
{{ end }}

{{define "method" }}
	{{ if .Fields -}}
		{{ .Description }}
		type {{ .Name }} struct { 
			{{range $field := .Fields -}}
				{{ $field.Name }} {{ $field.Type }} {{ $field.Tag }} // {{ $field.Description }} 
			{{end -}} 
		}
	{{ end }}

	{{ if or .InputFileFields .NestedInputFileFields }}
		func (x *{{ .Name }}) getFiles() map[string]*InputFile {
			media := map[string]*InputFile{}
	
			{{ range $field := .InputFileFields -}}
				if x.{{ $field.Name }} != nil {
					if x.{{ $field.Name }}.IsUploadable() {
            			media["{{ $field.FieldName }}"] = x.{{ $field.Name }}
          			}
				}
			{{ end }}

			{{ range $field := .NestedInputFileFields -}}
				{{ if isArray $field.Type -}} 
					for idx, m := range x.{{ $field.Name }} {
						for key, value := range m.getFiles() {
							media[fmt.Sprintf("%d.{{ $field.FieldName }}.%s", idx, key)] = value
						}            
					}
				{{ else -}}
					{{ if $field.IsOptional -}}
					if x.{{ $field.Name }} != nil { 
					{{ end -}}
					
						for key, value := range x.{{ $field.Name }}.getFiles() {
							media["{{ $field.FieldName }}."+key] = value
						}
					
					{{ if $field.IsOptional -}} 
					} 
					{{ end -}}
				{{ end -}}
			{{ end }}

			return media
		}

		func (x *{{ .Name }}) getParams() (map[string]string, error) {
			payload := map[string]string{}

			{{range $field := .Fields -}}
				{{ if $field.StringifyCode -}}
					{{ if not $field.IsOptional -}}
						{{ $field.StringifyCode }}
					{{ else -}} 
						if x.{{ $field.Name }} {{ if not (eq $field.DefaultValue "false") -}}!= {{ $field.DefaultValue }}{{ end -}} {
							{{ $field.StringifyCode }}
						} 
					{{ end -}}
				{{ end -}}
			{{ end }} 

			return payload, nil
		}
	{{ end }}

	{{ .Description }}
	func (api *API) {{ .Name }}({{ if .Fields -}}payload *{{ .Name }}{{ end -}}) ({{ .ReturnType }}, error) {
		{{ if or .InputFileFields .NestedInputFileFields -}}
			if files := payload.getFiles(); len(files) != 0 {
				params, err := payload.getParams()
				if err != nil {
				return {{ .EmptyReturnValue }}, err
				}
				return callMultipart[{{ .ReturnType }}](api, "{{.MethodName}}", params, files)
			}
		{{ end -}}
		return callJson[{{ .ReturnType }}](api, "{{.MethodName}}", {{ if .Fields -}}payload{{ else -}}nil{{ end -}})
	}
{{ end }}

{{ $impx := .Implementers }}
{{ $sections := .Sections }}

{{ template "header" }}
{{ range $section := $sections }}
	{{ if or $section.IsInterface $section.InterfaceOf }}
		{{ template "interface" (getInterfaceTemplate $section $sections) }}
	{{ else if isMethod .Name }}
		{{template "method" (getMethodTemplate $section $sections)}}
	{{ else }}
		{{ template "type" (getTypeTemplate $section $sections $impx) }}
	{{ end }}
{{ end }}