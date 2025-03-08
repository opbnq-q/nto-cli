package utils

func GetNotImplementedStructs(modelsFilePath string) []string {
	var models []string
	for _, model := range GetStructsList(modelsFilePath) {
		if !IsEntityImplemented(model) {
			models = append(models, model)
		}
	}
	return models
}
