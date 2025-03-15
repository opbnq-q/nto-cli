package generation

import (
	"github.com/opbnq-q/nto-cli/model"
	"github.com/opbnq-q/nto-cli/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Generate(model *model.Model) {
	mkPath := filepath.Join(utils.FindFrontendPath(), strings.ToLower(model.Name))
	if err := os.Mkdir(mkPath, 0755); err != nil {
		log.Fatalf("Failed to mkdir for model: %s", err)
	}
	GenerateService(model, mkPath)
	GenerateScheme(model, mkPath)
}
