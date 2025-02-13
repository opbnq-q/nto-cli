package main

import (
	"fmt"
	input "nto_cli/cmd"
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

	fmt.Printf("%+v\n", structFields)
}