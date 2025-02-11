package input

import "fmt"

func Input() (string, string) {
	var structName, path string
	fmt.Scan(&structName, &path)
	return structName, path
}