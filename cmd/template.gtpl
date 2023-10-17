// CODE GENERATED. DO NOT EDIT.
package tgo

import (
	"fmt"
	"strconv"
	"encoding/json"
)

{{ $impx := .Implementers }}
{{ $sections := .Sections }}

{{ range $section := $sections }}
	{{range $desc := $section.Description -}}// {{ $desc }}{{end -}}
	{{ if $section.IsInterface }} {{/* INTERFACE */}}
		type {{ upperFirstLetter $section.Name }} interface {	
			// Is{{ $section.Name }} does nothing and is only used to enforce type-safety
			Is{{ $section.Name }}()

			{{/* I had to do it somehow, so here's the result :skull: */}}
			{{ if eq (upperFirstLetter $section.Name) "InputMedia" -}}getFiles() map[string]*InputFile{{ end -}}
		}
	{{ else }} {{/* TYPE */}}
		{{ $isMethod := eq (index (lowerFirstLetter $section.Name) 0) (index $section.Name 0) -}}
		{{ if or $section.Fields (not $isMethod) -}}
			type {{ upperFirstLetter $section.Name }} struct { {{range $field := $section.Fields -}}
				{{ snakeToPascal $field.Name }} {{ getType $field.Name $field.Type $field.IsOptional }} {{ getTag $field.Name $field.IsOptional }} // {{ $field.Description }} 
			{{end -}} }
		{{ end }}

		{{ $implements := index $impx $section.Name }}
		{{ if $implements -}}
			func ({{ upperFirstLetter $section.Name }}) Is{{$implements}}() {}
		{{ end -}}

		{{ $inputFileFields := getMediaFields $section }}
		{{ $nestedInputFileFields := getNestedMediaFields $sections $section }}
		{{ if or $inputFileFields $nestedInputFileFields }}
			func (x *{{ upperFirstLetter $section.Name }}) getFiles() map[string]*InputFile {
				media := map[string]*InputFile{}
				
				{{ range $field := $inputFileFields -}} 
					{{ if $field.IsOptional -}} if x.{{ snakeToPascal $field.Name }} != nil { {{ end -}}
					if x.{{ snakeToPascal $field.Name }}.IsUploadable() {
						media["{{ $field.Name }}"] = x.{{ snakeToPascal $field.Name }}
					}
					{{ if $field.IsOptional -}} } {{ end -}}
				{{ end -}}

				{{ range $field := $nestedInputFileFields -}}
					{{ if isArray (getType $field.Name $field.Type $field.IsOptional) -}} 
						for idx, m := range x.{{ snakeToPascal $field.Name }} {
							for key, value := range m.getFiles() {
								value.Value = fmt.Sprintf("%d.{{ $field.Name }}.%s", idx, key)
								media[value.Value] = value
							}						
						}
					{{ else -}}
						{{ if $field.IsOptional -}}if x.{{ snakeToPascal $field.Name }} != nil { {{ end -}}
						for key, value := range x.{{ snakeToPascal $field.Name }}.getFiles() {
							value.Value = "{{ $field.Name }}."+key
							media[value.Value] = value
						}
						{{ if $field.IsOptional -}}}{{ end -}}
					{{ end -}}
				{{ end }}

				return media
			}
		{{ end }}

		{{ $canBeMedia := and $isMethod (or $inputFileFields $nestedInputFileFields) }}
		{{ if and $isMethod $canBeMedia }}
			func (x *{{ upperFirstLetter $section.Name }}) getParams() (map[string]string, error) {
				payload := map[string]string{}

				{{range $field := $section.Fields -}}
					{{ $stringerField := getStringerMethod "x" $field.Name $field.Type $field.IsOptional -}}
					{{ if not $field.IsOptional -}}
						{{ $stringerField }}
					{{ else -}} 
						{{ $dv := getDefaultValue (getType $field.Name $field.Type $field.IsOptional) -}}
						if x.{{snakeToPascal $field.Name }} {{ if not (eq $dv "false") -}}!= {{ $dv }}{{ end -}} {
							{{ $stringerField }}
						} 
					{{ end -}}
				{{ end -}} 

				return payload, nil
			}
		{{ end }}

		{{ if $isMethod }}
			{{ $returnType := getType $section.Name (extractReturnType $section.Description) true }}
			{{range $desc := $section.Description -}}// {{ $desc }}{{end}}
			func (api *API) {{upperFirstLetter $section.Name}}({{ if $section.Fields -}}payload *{{upperFirstLetter $section.Name}}{{ end -}}) ({{ $returnType }}, error) {
				{{ if $canBeMedia -}}
				if files := payload.getFiles(); len(files) != 0 {
					params, err := payload.getParams()
					if err != nil {
						return {{ getDefaultValue $returnType}}, err
					}
					return callMultipart[{{ $returnType }}](api, "{{$section.Name}}", params, files)
				}
				{{ end -}}
				return callJson[{{ $returnType }}](api, "{{$section.Name}}", {{ if $section.Fields -}}payload{{ else -}}nil{{ end -}})
			}
		{{ end }}
	{{ end }}
{{ end }}