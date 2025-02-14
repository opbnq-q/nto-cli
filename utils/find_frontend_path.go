package utils

import (
	"errors"
	"os"
	"strings"
)

func FindFrontendPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirs := strings.Split(currentPath, "\\")
	if dirs[len(dirs) - 2] + "/" + dirs[len(dirs) - 1] != "frontend/src" {
		panic(errors.New("You're not in frontend/src"))
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