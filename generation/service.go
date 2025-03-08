package generation

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/service.tmpl
var ServiceTemplate string

type ServiceTemplateContext struct {
	LowerModelName string
	ModelName      string
	ServicesPath   string
}

func GenerateService(structName, mkPath string) {
	servicePath := filepath.Join(mkPath, strings.ToLower(structName)+".service.ts")
	serviceFile, err := os.Create(servicePath)
	if err != nil {
		log.Fatalf("Failed to create service file: %s", err)
	}

	defer func(serviceFile *os.File) {
		err := serviceFile.Close()
		if err != nil {
			log.Fatalf("Failed to close file: %s", err)
		}
	}(serviceFile)

	context := ServiceTemplateContext{
		LowerModelName: strings.ToLower(structName),
		ModelName:      structName,
		ServicesPath:   GolangServicesPath,
	}

	serviceTemplate, err := template.New("service").Parse(ServiceTemplate)

	if err != nil {
		panic(fmt.Sprintf("Failed to parse service template: %s", err))
	}

	err = serviceTemplate.Execute(serviceFile, context)
	if err != nil {
		log.Fatalf("Failed to execute template: %s", err)
	}
	log.Printf("Service for `%s` model is written: %s", structName, servicePath)
}
