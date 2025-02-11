package utils

import "strings"

func ContainsMany(str string, substrs... string) bool {
	count := 0
	for _, substr := range substrs {
		if strings.Contains(str, substr) {
			count++;
		} else {
			return false;
		}
	}
	return count == len(substrs)
}