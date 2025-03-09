package main

import (
	"log"
	"nto_cli/cmd"
	"nto_cli/generation"
	"nto_cli/model"
	"nto_cli/utils"
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
