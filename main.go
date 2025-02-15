package main

import (
	"fmt"
	input "nto_cli/cmd"
	// "nto_cli/generation"
	"nto_cli/utils"
	"os"
)

func main() {
	structName, path := input.Input()
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	structFields := utils.GetStructFields(file, structName)
	for _, field := range structFields {
		fmt.Println(field.Generate())
	}

	// generation.Generate(structName, structFields)
}
