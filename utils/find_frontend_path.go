package utils

import (
	"errors"
	"os"
	"slices"
	"strings"
)

func FindFrontendPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirs := strings.Split(currentPath, "\\")
	if !slices.Contains(dirs, "frontend") {
		panic(errors.New("Frontend dir doesn't exist"))
	}
	var path string
	for i, dir := range dirs {
		if dir == "frontend" {
			break
		}
		if i > 0 {
			dir = "/" + dir
		}
		path += dir
	}
	return path
}