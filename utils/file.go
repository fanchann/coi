package utils

import "strings"

func GetFileLocation(fileName string) string {
	lastIndex := strings.LastIndex(fileName, "/")
	if lastIndex == -1 {
		return ""
	}
	extracted := fileName[:lastIndex]
	return extracted
}
