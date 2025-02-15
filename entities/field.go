package entities

import (
	"fmt"
	"slices"
	"strings"
)


type Field struct {
	Name     string
	Type     string
	Medatada []Medatada
}

var PRIMITIVE_TYPES = []string{"date", "number", "string", "multiple", "boolean"}

func (f *Field) GenerateType() string {
	result := "  type: {\n"

	if slices.Contains(PRIMITIVE_TYPES, strings.ToLower(f.Type)) {
		result += fmt.Sprintf(`    primitive: "%s",`, strings.ToLower(f.Type))
	} else {
		var field string
		for _, meta := range f.Medatada {
			if meta.Name == "field" {
				field = "['" + strings.Join(meta.Values, "', '") + "']"
			}
		}
		result += fmt.Sprintf(`    nested: {
      values: [],
      field: %s
    }, `, field)
}
result += "\n  }"
	for _, meta := range f.Medatada {
		if meta.Name == "many" {
			result += "\n many: true }"
		}
	}
	return result
}

func (f *Field) Generate() string {
	result := "{\n"
	for _, meta := range f.Medatada {
		if (meta.Name == "hidden") {
			result += "  hidden: true,\n"
		} else if meta.Name == "label" {
			result += fmt.Sprintf(`  russian: "%s",` + "\n", meta.Values[0])
		} else if (meta.Name == "readonly") {
			result += "  readonly: true,\n"
		}
	}
	result += f.GenerateType()
	return result + "\n}"
}