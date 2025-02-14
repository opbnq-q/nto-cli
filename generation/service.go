package generation

import (
	"fmt"
	"os"
	"strings"
)

func GetServiceBindPath(structName string) string {
	path := fmt.Sprintf("../bindings/app/internal/services/%sservice.ts", strings.ToLower(structName))
	return path
}

func GetServiceStructType(structName string) string {
	path := "../bindings/app/internal/services/models.ts"
	return path
}

func GetServiceType() string {
	path := "./types/service.type.ts"
	return path
}

func GenerateService(structName, mkPath string) {
	serviceFile, err := os.Create(mkPath + "/" + strings.ToLower(structName) + ".service.ts")
	if err != nil {
		panic(err)
	}
	defer serviceFile.Close()
	_, err = serviceFile.WriteString(
		fmt.Sprintf(`import { GetAll, Create, Delete, ExportToExcel, GetById, Update, Count } from "%s"
import type { %s } from "%s"
import type { Service } from "%s"
	
export class %sService implements Service {
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
`, GetServiceBindPath(structName), structName, GetServiceStructType(structName), GetServiceType(), structName, structName, structName, structName))
	if err != nil {
		panic(err)
	}
}
