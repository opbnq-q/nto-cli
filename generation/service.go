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
		`import { GetAll, Create, Delete, ExportToExcel, GetById, Update, Count } from "%s"
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
		return await Create(item)
	}

	async delete(item: %s) {
		return await Delete(item)
	}	
	async update(item: %s) {
		return await Update(item)
	}
	async count() {
		return await Count()
	}
	async exportToExcel() {
		return await ExportToExcel()
	}
}
`, utils.GetServiceBindPath(structName), structName, utils.GetServiceStructType(structName), utils.GetServiceType(), structName, structName, structName, structName, structName))
	if err != nil {
		panic(err)
	}
}
