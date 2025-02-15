package entities

import (
	"strings"
)

type Medatada struct {
	Name   string
	Values []string
}

func NewMetadata(tag string) *Medatada {
	tag = strings.TrimSpace(tag)
	tagName := ""
	var values []string
	if strings.Contains(tag, "=") {
		tagName = tag[:strings.Index(tag, "=")]
		if tag[strings.Index(tag, "=") + 1] == '[' {
			values = append(values, strings.Split(tag[strings.Index(tag, "=") + 2:len(tag)-1], ";")...)
			for i := range values {
				values[i] = strings.TrimSpace(values[i])
			}
		} else {
			values = append(values, strings.TrimSpace(tag[strings.Index(tag, "=") + 1:]))
		}
	} else {
		tagName = tag
	}
	return &Medatada{
		Name: strings.TrimSpace(strings.ToLower(tagName)),
		Values: values,
	}
}