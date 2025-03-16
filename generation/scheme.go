package generation

import (
	_ "embed"
	"fmt"
	"github.com/opbnq-q/nto-cli/model"
	"github.com/opbnq-q/nto-cli/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/scheme.tmpl
var SchemeTemplate string

type Dependency struct {
	ImportName  string
	ServiceName string
	LowerName   string
	FieldName   string
}

type TemplateData struct {
	StructName   string
	LowerName    string
	Fields       []model.Field
	Dependencies []Dependency
}

func GenerateScheme(model *model.Model, mkPath string) {
	data := TemplateData{
		StructName:   model.Name,
		LowerName:    strings.ToLower(model.Name),
		Fields:       model.Fields,
		Dependencies: processDependencies(model.Fields),
	}

	schemeFilename := strings.ToUpper(model.Name[:1]) + strings.ToLower(model.Name[1:]) + "Scheme.vue"
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
	log.Printf("Scheme for `%s` model is written: %s", model.Name, schemeFilePath)
	_ = utils.FormatFilesWithPrettier([]string{schemeFilePath})
}

func processDependencies(fields []model.Field) []Dependency {
	var dependencies []Dependency
	encountered := make(map[string]bool)

	for _, field := range fields {
		dep := field.Metadata.RelatedModel
		if dep == "" || encountered[dep] {
			continue
		}
		encountered[dep] = true

		dependencies = append(dependencies, Dependency{
			ImportName:  strings.ToUpper(dep[:1]) + strings.ToLower(dep[1:]) + "Service",
			ServiceName: strings.ToLower(dep) + "Service",
			LowerName:   strings.ToLower(dep),
			FieldName:   field.Name,
		})

	}

	return dependencies
}
