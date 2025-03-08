package entities

import (
	"bytes"
	"fmt"
	"text/template"
)

const fieldTemplate = `{
{{ range .Metadata }}{{ if eq .Name "hidden" }}  hidden: true,
{{ else if eq .Name "label" }}  russian: "{{ index .Values 0 }}",
{{ else if eq .Name "readonly" }}  readonly: true,
{{ end }}{{ end }}{{ if .IsArray }}  many: true,
{{ end }}{{ .GeneratedType }}
}`

type FieldTemplateContext struct {
	Metadata      []Metadata
	IsArray       bool
	GeneratedType string
}

func (f *Field) Generate() string {
	tmpl, err := template.New("field").Parse(fieldTemplate)
	if err != nil {
		panic(fmt.Sprintf("Error parsing field template: %v", err))
	}

	isArray := len(f.Type) >= 2 && f.Type[0:2] == "[]"

	data := FieldTemplateContext{
		Metadata:      f.Metadata,
		IsArray:       isArray,
		GeneratedType: f.GenerateType(),
	}

	var result bytes.Buffer
	if err := tmpl.Execute(&result, data); err != nil {
		panic(fmt.Sprintf("Error executing field template: %v", err))
	}

	return result.String()
}
