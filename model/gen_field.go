package model

import (
	"bytes"
	"fmt"
	"text/template"
)

const fieldTemplate = `{
{{ if .Metadata.Hidden }}  hidden: true,
{{ end }}{{ if .Metadata.Label }}  russian: "{{ .Metadata.Label }}",
{{ end }}{{ if .Metadata.Readonly }}  readonly: true,
{{ end }}{{ if .IsArray }}  many: true,
{{ end }}{{ if eq .Datatype "datetime" }} date: true,
{{ end }}{{ .GeneratedType }}
}`

type FieldTemplateContext struct {
	Metadata      FieldMetadata
	IsArray       bool
	GeneratedType string
	Datatype      string
}

func (field *Field) GenerateFieldCode() string {
	tmpl, err := template.New("field").Parse(fieldTemplate)
	if err != nil {
		panic(fmt.Sprintf("Error parsing field template: %v", err))
	}

	data := FieldTemplateContext{
		Metadata:      field.Metadata,
		IsArray:       field.Metadata.IsSlice,
		GeneratedType: field.GenerateType(),
		Datatype:      field.Metadata.Datatype,
	}

	var result bytes.Buffer
	if err := tmpl.Execute(&result, data); err != nil {
		panic(fmt.Sprintf("Error executing field template: %v", err))
	}

	return result.String()
}
