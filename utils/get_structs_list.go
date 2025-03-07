package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetStructList(filePath string) []string {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Fatalf("Failed to open a file: %s", err)
	}

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

	if s.Err() != nil {
		log.Fatalf("Unexpected scanner error: %s", err)
	}

	return structNames
}
