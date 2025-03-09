package entities

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
	"text/template"
)

var PrimitiveTypes = map[string]string{
	"date":    "date",
	"number":  "number",
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

func (f *Field) GenerateType() string {
	keys := make([]string, 0, len(PrimitiveTypes))
	for k := range PrimitiveTypes {
		keys = append(keys, k)
	}

	var data TypeTemplateContext

	if slices.Contains(keys, strings.ToLower(f.Type)) {
		data.IsPrimitive = true
		data.PrimitiveType = PrimitiveTypes[strings.ToLower(f.Type)]
	} else {
		data.IsPrimitive = false
		field := "[]"
		for _, meta := range f.Metadata {
			if meta.Name == "field" && len(meta.Values) > 0 {
				field = "['" + strings.Join(meta.Values, "', '") + "']"
				break
			}
		}
		data.Field = field
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
