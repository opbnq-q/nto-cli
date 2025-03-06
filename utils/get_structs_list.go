package utils

import (
    "bufio"
    "os"
    "strings"
)

func GetStructList(filePath string) ([]string) {
    file, err := os.Open(filePath)
    if err != nil {
		panic(err)
    }
    defer file.Close()
    var structNames []string
    s := bufio.NewScanner(file)
    for s.Scan() {
        line := s.Text()
        if strings.Contains(line, "type ") && strings.Contains(line, " struct") {
            start := strings.Index(line, "type ") + 5
            end := strings.Index(line, " struct")
            name := strings.TrimSpace(line[start:end])
            if name != "" {
                structNames = append(structNames, name)
            }
        }
    }
    if err := s.Err(); err != nil {
        panic(err)
    }
    return structNames
}