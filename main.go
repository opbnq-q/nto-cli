package main

import (
	"github.com/opbnq-q/nto-cli/cmd"
	"github.com/opbnq-q/nto-cli/generation"
	"github.com/opbnq-q/nto-cli/model"
	"github.com/opbnq-q/nto-cli/utils"
	"log"
)

func main() {
	log.SetFlags(0)
	modelsPath := utils.GetModelsPath()
	models, err := model.ParseModelsPackage(modelsPath)
	if err != nil {
		log.Fatalf("Failed to parse models: %s", err)
	}
	selectedModels := cmd.SelectionInput(models)

	for _, m := range *selectedModels {
		generation.Generate(&m)
	}
}
