package utils

import (
	"nto_cli/entities"
	"strings"

	"github.com/fatih/structtag"
)

func SplitStructField(field string) *entities.Field {
	if strings.Contains(field, "type") {
		return nil
	}

	startBacktip := strings.Index(field, "`")
	endBacktip := -1
	var metadata []entities.Medatada
	if startBacktip > -1 {
		endBacktip = strings.Index(field[startBacktip + 1:], "`")
		if endBacktip > -1 {
			endBacktip += startBacktip + 1
			meta := field[startBacktip + 1 : endBacktip]
			tags, err := structtag.Parse(meta)
			if err != nil {
				panic(err)
			}
			uiTag, err := tags.Get("ui")
			if err != nil {
				panic(err)
			}
			uiTags := append([]string{uiTag.Name}, uiTag.Options...)
			for _, t := range uiTags {
				analyzed := entities.NewMetadata(t)
				if analyzed != nil {
					metadata = append(metadata, *analyzed)
				}
			}
		}
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

	return &entities.Field{
		Medatada: metadata,
		Type: dataType,
		Name: name,
	}
}