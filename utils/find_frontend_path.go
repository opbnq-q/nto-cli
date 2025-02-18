package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func FindFrontendPath() string {
    currentPath, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    var dirs []string
    for currentPath != "/" {
        dir, file := filepath.Split(currentPath)
        if file != "" {
            dirs = append([]string{file}, dirs...)
        }
        currentPath = filepath.Clean(dir)
    }

    if len(dirs) < 2 || dirs[len(dirs)-2] + "/" + dirs[len(dirs)-1] != "frontend/src" {
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

    if dirs[0] == "home" {
      path = "/" + path  
    }
    return path
}