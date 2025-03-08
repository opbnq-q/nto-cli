package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FindFrontendPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to find path for frontend: %s", err)
	}

	frontendSuffix := filepath.Join("frontend")
	frontendSourceSuffix := filepath.Join("frontend", "src")
	if strings.HasSuffix(currentPath, frontendSourceSuffix) {
		return currentPath
	}

	var path string

	if strings.HasSuffix(currentPath, frontendSuffix) {
		path = filepath.Join(currentPath, "src")
	} else {
		path = filepath.Join(currentPath, frontendSourceSuffix)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("Frontend source directory not found at %s", path)
	}

	return path
}
