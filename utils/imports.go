package utils

import (
	"fmt"
	"strings"
)

func GetServiceBindPath(structName string) string {
	path := fmt.Sprintf("../../bindings/app/internal/services/%sservice.ts", strings.ToLower(structName))
	return path
}

func GetServiceStructType(structName string) string {
	path := "../../bindings/app/internal/services/models.ts"
	return path
}

func GetServiceType() string {
	path := "../types/service.type.ts"
	return path
}
