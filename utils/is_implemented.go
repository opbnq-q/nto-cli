package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func IsEntityImplemented(entityName string) bool {
	entityDirectory := filepath.Join(FindFrontendPath(), strings.ToLower(entityName))
	if _, err := os.Stat(entityDirectory); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
