package utils

import "strings"

func SplitBySingleSpace(input string) []string {
    parts := strings.Split(strings.TrimSpace(input), " ")
    var result []string
    for _, p := range parts {
        if p != "" {
            result = append(result, p)
        }
    }
    return result
}