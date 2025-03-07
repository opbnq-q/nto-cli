package entities

import (
	"fmt"
	"slices"
	"strings"
)

type Field struct {
	Name     string
	Type     string
	Medatada []Metadata
}

var PRIMITIVE_TYPES = map[string]string{
	"date":     "date",
	"number":   "number",
	"string":   "string",
	"multiple": "multiple",
	"boolean":  "boolean",
	"bool":     "boolean",
	"int":      "number",
	"uint":     "number",
	"float32":  "number",
	"float64":  "number",
	"int32":    "number",
	"int64":    "number",
	"uint32":   "number",
	"uint64":   "number",
	"int8":     "number",
	"int16":    "number",
	"uint8":    "number",
	"uint16":   "number",
	"byte":     "number",
	"rune":     "number",
}

func (f *Field) GenerateType() string {
	result := "  type: {\n"

	keys := make([]string, 0, len(PRIMITIVE_TYPES))
	for k := range PRIMITIVE_TYPES {
		keys = append(keys, k)
	}

	if slices.Contains(keys, strings.ToLower(f.Type)) {
		result += fmt.Sprintf(`    primitive: "%s",`, PRIMITIVE_TYPES[strings.ToLower(f.Type)])
	} else {
		var field string = "[]"
		for _, meta := range f.Medatada {
			if meta.Name == "field" {
				if len(meta.Values) > 0 {
					field = "['" + strings.Join(meta.Values, "', '") + "']"
				}
			}
		}
		result += fmt.Sprintf(`    nested: {
      values: [],
      field: %s
    }, `, field)
	}
	result += "\n  },"
	return result
}

func (f *Field) Generate() string {
	result := "{\n"
	for _, meta := range f.Medatada {
		if meta.Name == "hidden" {
			result += "  hidden: true,\n"
		} else if meta.Name == "label" {
			result += fmt.Sprintf(`  russian: "%s",`+"\n", meta.Values[0])
		} else if meta.Name == "readonly" {
			result += "  readonly: true,\n"
		}
	}
	if f.Type[0:2] == "[]" {
		result += "  many: true,\n"
	}
	result += f.GenerateType()
	return result + "\n}"
}
