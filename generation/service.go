package generation

import (
	"fmt"
	"os"
)

func GenerateService(structName, mkName string) {
	serviceFile, err := os.Create(mkName + "/service.ts") 
	if err != nil {
		panic(err)
	}
	defer serviceFile.Close()
	_, err = serviceFile.WriteString(fmt.Sprintf(`
export class %sService {
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