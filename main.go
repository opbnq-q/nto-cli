package main

import (
	"nto_cli/cmd"
	"nto_cli/generation"
	"nto_cli/utils"
	"os"
)

func main() {
	structNames, path := cmd.SelectionInput()

	for _, structName := range structNames {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		structFields := utils.GetStructFields(file, structName)
		file.Close()
		generation.Generate(structName, structFields)
	}
}
