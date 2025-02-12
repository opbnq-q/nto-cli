package input

import "fmt"

func Input() (string, string) {
	fmt.Print("struct name, path to file (including struct): ")
	var structName, path string
	fmt.Scan(&structName, &path)
	return structName, path
}