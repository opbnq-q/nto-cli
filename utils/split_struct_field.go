package utils

import (
	"errors"
	"log"
	"nto_cli/entities"
	"strings"

	"github.com/fatih/structtag"
)

func SplitStructField(field string) (*entities.Field, error) {
	if strings.Contains(field, "type") {
		return nil, nil
	}
	if len(strings.TrimSpace(field)) < 2 {
		return nil, errors.New("unexpected end of struct field")
	}
	startBacktick := strings.Index(field, "`")
	endBacktick := -1

	var metadata []entities.Metadata

	if startBacktick > -1 {
		endBacktick = strings.Index(field[startBacktick+1:], "`")
		if endBacktick > -1 {
			endBacktick += startBacktick + 1
			meta := field[startBacktick+1 : endBacktick]
			tags, err := structtag.Parse(meta)
			if err != nil {
				log.Fatalf("failed to parse struct tag: %s", err)
			}
			uiTag, err := tags.Get("ui")

			if err == nil {
				uiTags := append([]string{uiTag.Name}, uiTag.Options...)
				for _, t := range uiTags {
					analyzed := entities.NewMetadata(t)
					if analyzed != nil {
						metadata = append(metadata, *analyzed)
					}
				}
			}
		}
	} else {
		startBacktick = len(field)
	}

	field = strings.TrimSpace(field[:startBacktick])

	data := SplitBySingleSpace(field)

	name := data[0]
	dataType := data[1]
	return &entities.Field{
		Metadata: metadata,
		Type:     dataType,
		Name:     name,
	}, nil
}
