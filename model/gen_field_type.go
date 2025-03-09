package model

import (
	"bytes"
	"fmt"
	"go/ast"
	"strings"
	"text/template"
)

var PrimitiveTypes = map[string]string{
	"date":    "date",
	"string":  "string",
	"boolean": "boolean",
	"bool":    "boolean",
	"int":     "number",
	"uint":    "number",
	"float32": "number",
	"float64": "number",
	"int32":   "number",
	"int64":   "number",
	"uint32":  "number",
	"uint64":  "number",
	"int8":    "number",
	"int16":   "number",
	"uint8":   "number",
	"uint16":  "number",
	"byte":    "number",
	"rune":    "number",
}

const typeTemplate = `  type: {
{{ if .IsPrimitive }}    primitive: "{{ .PrimitiveType }}",{{ else }}    nested: {
     values: [],
     field: {{ .Field }}
   }, {{ end }}
 },`

type TypeTemplateContext struct {
	IsPrimitive   bool
	PrimitiveType string
	Field         string
}

func (field *Field) GenerateType() string {
	keys := make([]string, 0, len(PrimitiveTypes))
	for k := range PrimitiveTypes {
		keys = append(keys, k)
	}

	var data TypeTemplateContext

	if field.Metadata.IsPrimitiveType {
		data.IsPrimitive = true
		var typeName string
		// TODO: resolve datatype in other function
		switch field.Metadata.Datatype {
		case "datetime":
			typeName = "date"
		default:
			typeName = field.Type.(*ast.Ident).Name
		}
		data.PrimitiveType = PrimitiveTypes[typeName]
	} else {
		data.IsPrimitive = false
		data.Field = "['" + strings.Join(field.Metadata.RelatedFields, "', '") + "']"
	}

	tmpl, err := template.New("type").Parse(typeTemplate)
	if err != nil {
		panic(fmt.Sprintf("Error parsing template: %v", err))
	}

	var result bytes.Buffer
	if err := tmpl.Execute(&result, data); err != nil {
		panic(fmt.Sprintf("Error executing template: %v", err))
	}

	return result.String()
}
