package generation

import (
	"fmt"
	"nto_cli/entities"
	"nto_cli/utils"
	"os"
	"strings"
)

func GenerateScheme(structName string, fields []entities.Field, mkPath string) {
	schemeFile, err := os.Create(mkPath + "/" + strings.ToUpper(structName[:1]) + strings.ToLower(structName[1:]) + "Scheme.vue")
	if err != nil {
		panic(err)
	}
	defer schemeFile.Close()
	_, err = schemeFile.WriteString(fmt.Sprintf(
		`<script setup lang="ts">
import Table from '../table/Table.vue'
import { onMounted, reactive } from 'vue'
import { getDefaultValues } from '../utils/structs/defaults.util'
import S from './%s.service.ts'
import type { Scheme } from '../types/scheme.type'
import { %s } from '%s'

const service = new S

%s

const scheme: Scheme<%s> = reactive(%s)

const getDefaults = ()  => getDefaultValues(scheme)

</script>

<template>
	<main class="w-screen h-screen">
		<Table :scheme :service :getDefaults></Table>
	</main>
</template>
`, strings.ToLower(structName), structName, utils.GetServiceStructType(structName), LoadDependencies(fields), structName, GenerateFields(fields)))
	if err != nil {
		panic(err)
	}
}

func GenerateFields(fields []entities.Field) string {
	result := "{\n"
	for _, field := range fields {
		result += field.Name + ":" + field.Generate() + ", \n"
	}
	return result + "\n}"
}

func LoadDependencies(fields []entities.Field) string {
	type Dependency struct{
		fieldName string
		dependencyName string
	} 

	result := ""
	dependencies := []Dependency{}
	for _, field := range fields {
		for _, meta := range field.Medatada {
			if meta.Name == "data" {
				dependency := meta.Values[0]
				dependencies = append(dependencies, Dependency{
					fieldName: field.Name,
					dependencyName: dependency,
				})
				result += fmt.Sprintf("import %sService from '../%s/%s.service.ts'\n", dependency, strings.ToLower(dependency), strings.ToLower(dependency))
				result += fmt.Sprintf("const %sService = new %sService\n", strings.ToLower(dependency), strings.ToUpper(dependency[:1]) + strings.ToLower(dependency[1:]))
			}
		}
	}
	insertIntoScheme := ""
	for _, dep := range dependencies {
		insertIntoScheme += fmt.Sprintf("(scheme as any).%s.type!.nested!.values = await %sService.readAll()\n", dep.fieldName, strings.ToLower(dep.dependencyName))
	}

	result += fmt.Sprintf(`onMounted(async () => {
  	%s
})` + "\n", insertIntoScheme)
	return result
}
