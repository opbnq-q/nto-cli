package generation

import (
	"fmt"
	"nto_cli/entities"
	"os"
	"strings"
)

func Generate(structName string, fields []entities.Field) {
	mkName := strings.ToLower(fmt.Sprintf("./%s", structName))
	if err := os.Mkdir(mkName, 0755); err != nil {
		panic(err)
	}
	GenerateService(structName, mkName)
}