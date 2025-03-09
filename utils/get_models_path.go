package utils

import (
	"log"
	"os"
)

func GetModelsPath() string {
	if len(os.Args) == 1 {
		log.Fatalf("Please provide path to models.go")
	}

	return os.Args[1]
}
