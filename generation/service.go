package generation

import (
	"fmt"
	"nto_cli/utils"
	"os"
	"strings"
)



func GenerateService(structName, mkPath string) {
	serviceFile, err := os.Create(mkPath + "/" + strings.ToLower(structName) + ".service.ts")
	if err != nil {
		panic(err)
	}
	defer serviceFile.Close()
	_, err = serviceFile.WriteString(fmt.Sprintf(
		`import { GetAll, Create, Delete, GetById, Update, Count } from "%s"
import type { %s } from "%s"
import type { IService } from "%s"
	
export default class %sService implements IService<%s> {
	async read(id: number) {
		return await GetById(id)
	}

	async readAll() {
		return await GetAll()
	}

	async create(item: %s) {
		await Create(item)
	}

	async delete(id: number) {
		return await Delete(id)
	}	
	async update(item: %s) {
		await Update(item)
	}
	async count() {
		return await Count()
	}
}
`, utils.GetServiceBindPath(structName), structName, utils.GetServiceStructType(structName), utils.GetServiceType(), structName, structName, structName, structName))
	if err != nil {
		panic(err)
	}
}
