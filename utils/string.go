package utils

import (
	"strings"
	"unicode"
)

func MakePrettyFileName(input string) string {
	var result string
	checkInputStr := removeI(input)

	for _, char := range checkInputStr {
		if unicode.IsUpper(char) {
			if result != "" {
				result += "_"
			}
			result += string(unicode.ToLower(char))
		} else {
			result += string(char)
		}
	}

	return result
}

func removeI(s string) string {
	result := strings.Replace(s, "I", "", -1)
	return result
}
