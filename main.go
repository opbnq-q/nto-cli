package main

import (
	"log"
	"nto_cli/cmd"
	"nto_cli/generation"
	"nto_cli/utils"
	"os"
)

func main() {
	log.SetFlags(0)
	structNames, path := cmd.SelectionInput()

	for _, structName := range structNames {
		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("Failed to open file: %s", err)
		}
		structFields := utils.GetStructFields(file, structName)
		_ = file.Close()
		generation.Generate(structName, structFields)
	}
}
