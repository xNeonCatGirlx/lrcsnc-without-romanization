package util

import "regexp"

var badChars = regexp.MustCompile(`[:;!?,\.\[\]<>\/\\*|]+`)

func RemoveBadCharacters(str string) string {
	return badChars.ReplaceAllString(str, "_")
}
