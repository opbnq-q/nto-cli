package generation

import (
	"fmt"
	"nto_cli/entities"
	"nto_cli/utils"
	"os"
	"strings"
)

func Generate(structName string, fields []entities.Field) {
	mkPath := strings.ToLower(fmt.Sprintf("%s/frontend/src/%s", utils.FindFrontendPath() , structName))
	if err := os.Mkdir(mkPath, 0755); err != nil {
		panic(err)
	}
	GenerateService(structName, mkPath)
}