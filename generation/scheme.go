package generation

import (
	_ "embed"
	"fmt"
	"log"
	"nto_cli/entities"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/scheme.tmpl
var SchemeTemplate string

const GolangServicesPath = "../../bindings/app/internal/services"

type Dependency struct {
	ImportName  string
	ServiceName string
	LowerName   string
	FieldName   string
}

type TemplateData struct {
	StructName         string
	LowerName          string
	GolangServicesPath string
	Fields             []entities.Field
	Dependencies       []Dependency
}

func GenerateScheme(structName string, fields []entities.Field, mkPath string) {
	data := TemplateData{
		StructName:         structName,
		LowerName:          strings.ToLower(structName),
		GolangServicesPath: GolangServicesPath,
		Fields:             fields,
		Dependencies:       processDependencies(fields),
	}

	schemeFilename := strings.ToUpper(structName[:1]) + strings.ToLower(structName[1:]) + "Scheme.vue"
	schemeFilePath := filepath.Join(mkPath, schemeFilename)
	schemeFile, err := os.Create(schemeFilePath)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer func(schemeFile *os.File) {
		err := schemeFile.Close()
		if err != nil {
			log.Fatalf("Failed to close file: %s", err)
		}
	}(schemeFile)

	tmpl, err := template.New("scheme").Parse(SchemeTemplate)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse template: %s", err))
	}

	err = tmpl.Execute(schemeFile, data)
	if err != nil {
		log.Fatalf("Failed to execute template: %s", err)
	}
	log.Printf("Scheme for `%s` model is written: %s", structName, schemeFilePath)
}

func processDependencies(fields []entities.Field) []Dependency {
	var dependencies []Dependency

	for _, field := range fields {
		for _, meta := range field.Metadata {
			if meta.Name == "data" {
				dependency := meta.Values[0]
				dependencies = append(dependencies, Dependency{
					ImportName:  strings.ToUpper(dependency[:1]) + strings.ToLower(dependency[1:]) + "Service",
					ServiceName: strings.ToLower(dependency) + "Service",
					LowerName:   strings.ToLower(dependency),
					FieldName:   field.Name,
				})
			}
		}
	}

	return dependencies
}
