package utils

import "strings"

func ContainsMany(str string, substrings ...string) bool {
	var matches int
	for _, substr := range substrings {
		if strings.Contains(str, substr) {
			matches++
		} else {
			return false
		}
	}
	return matches == len(substrings)
}
