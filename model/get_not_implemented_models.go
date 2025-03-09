package model

import "nto_cli/utils"

func GetNotImplementedModels(models []Model) []Model {
	var unimplementedModels []Model
	for _, m := range models {
		if !utils.IsEntityImplemented(m.Name) {
			unimplementedModels = append(unimplementedModels, m)
		}
	}
	return unimplementedModels
}
