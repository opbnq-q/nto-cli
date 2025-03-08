package utils

import (
	"fmt"
	"log"
	"os/exec"
)

func FormatFilesWithPrettier(files []string) error {
	if len(files) == 0 {
		return fmt.Errorf("empty file list")
	}

	args := append([]string{"prettier", "--write"}, files...)

	cmd := exec.Command("npx", args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("formatter error: %s\nOutput: %s", err, output)
	}

	log.Println("prettier was applied")

	return nil
}
