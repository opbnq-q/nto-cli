package generation

import (
	"log"
	"nto_cli/entities"
	"nto_cli/utils"
	"os"
	"path/filepath"
	"strings"
)

func Generate(structName string, fields []entities.Field) {
	mkPath := filepath.Join(utils.FindFrontendPath(), strings.ToLower(structName))
	if err := os.Mkdir(mkPath, 0755); err != nil {
		log.Fatalf("Failed to mkdir for model: %s", err)
	}
	GenerateService(structName, mkPath)
	GenerateScheme(structName, fields, mkPath)
}
