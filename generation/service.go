package generation

import (
	"fmt"
	"nto_cli/utils"
	"os"
	"strings"
)

func GetServiceBindPath(structName string) string {
	path := utils.FindFrontendPath()
	path += fmt.Sprintf("/bindings/app/internal/services/%sservice.ts", strings.ToLower(structName))
	return path
}

func GenerateService(structName, mkPath string) {
	serviceFile, err := os.Create(mkPath + "/" + strings.ToLower(structName) + ".service.ts") 
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