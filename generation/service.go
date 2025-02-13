package generation

import (
	"fmt"
	"os"
	"strings"
)

func GenerateService(structName, mkName string) {
	serviceFile, err := os.Create(mkName + "/" + strings.ToLower(structName) + ".service.ts") 
	if err != nil {
		panic(err)
	}
	defer serviceFile.Close()
	_, err = serviceFile.WriteString(fmt.Sprintf(`export class %sService {
	async read() {

	}

	async readAll() {

	}

	async create() {
	
	}

	async delete() {
	
	}	
	async update() {
	
	}
}
`, structName))
	if err != nil {
		panic(err)
	}
}