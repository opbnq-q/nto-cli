package utils

import (
	"nto_cli/types"
	"strings"
)

func SplitStructField(field string) *types.Field {
	if strings.Contains(field, "type") {
		return nil
	}

	startBacktip := strings.Index(field, "`")
	var metadata []string
	if startBacktip > -1 {
		metadata = []string{field[startBacktip:]}
	} else {
		startBacktip = len(field)
	}

	field = strings.TrimSpace(field[:startBacktip])

	data := strings.Split(field, " ")

	if len(data) < 2 {
        return nil
    }

	name := strings.TrimSpace(data[0])

	dataType := strings.TrimSpace(data[1])

	return &types.Field{
		Medatada: metadata,
		Type: dataType,
		Name: name,
	}
}