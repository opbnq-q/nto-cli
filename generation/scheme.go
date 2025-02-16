package generation

import (
	"fmt"
	"nto_cli/entities"
	"nto_cli/utils"
	"os"
	"strings"
)

func GenerateScheme(structName string, fields []entities.Field, mkPath string) {
	schemeFile, err := os.Create(mkPath + "/" + strings.ToTitle(structName) + "Scheme.vue")
	if err != nil {
		panic(err)
	}
	defer schemeFile.Close()
	_, err = schemeFile.WriteString(fmt.Sprintf(
		`<script setup lang="ts">
import Table from ../table/Table.vue
import { onMounted, reactive } from 'vue'
import { getDefaultValues } from '../utils/structs/defaults.util'
import S from './%s.service.ts'
import type { Scheme } from '../types/scheme.type'
import { %s } from '%s'

const service = new S

onMounted(async () => {
	
})

const scheme: Scheme = reactive({
%s
})

getDefaults = ()  => getDefaultValues(scheme)

</script>

<template>
	<main class="w-screen h-screen>
		<Table :scheme :service :getDefaults></Table>
	</main>
</template>
`, strings.ToLower(structName), structName, utils.GetServiceStructType(structName), GenerateFields(fields)))
	if err != nil {
		panic(err)
	}
}

func GenerateFields(fields []entities.Field) string {
	result := "{\n"
	for _, field := range fields {
		result += field.Generate() + ", \n"
	}
	return result + "\n}"
}