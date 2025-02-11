package utils

import (
	"bufio"
	"nto_cli/types"
	"os"
	"strings"
)

func GetStructFields(file *os.File, structName string) []types.Field {
	bracketsCount := 1

	structFound := false

	structFields := []types.Field{}

	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan() && bracketsCount > 0; i++ {
		line := scanner.Text()
		if ContainsMany(line, structName, "type", "struct") {
			structFound = true
		}
		if structFound {
			bracketsCount += strings.Count(line, "{")
			bracketsCount -= strings.Count(line, "}")
			line = strings.TrimSpace(line)
			newField := SplitStructField(line)
			if newField != nil {
				structFields = append(structFields, *newField)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return structFields
}