package caseconvert

import (
	"regexp"
	"strings"
)

func ToCamel(input string) string {
	// Convert to camelCase using regex
	re := regexp.MustCompile(`(?:^|_)([a-zA-Z])`)
	camelCaseStr := re.ReplaceAllStringFunc(input, func(s string) string {
		return strings.ToUpper(string(s[0]))
	})
	return strings.ToLower(camelCaseStr)
}

func PascalToCamel(input string) string {
	re := regexp.MustCompile(`^([A-Z]+)`)
	camelCaseStr := re.ReplaceAllStringFunc(input, func(s string) string {
		return string(s[0]) + string(s[0]-'A'+'a')
	})
	return camelCaseStr
}

func SnakeToCamel(input string) string {
	re := regexp.MustCompile(`_([a-z0-9])`)
	camelCaseStr := re.ReplaceAllStringFunc(input, func(s string) string {
		return string(s[0] - 'a' + 'A')
	})
	return camelCaseStr
}

func ConstantToCamel(input string) string {
	re := regexp.MustCompile(`_([a-zA-Z0-9])`)
	camelCaseStr := re.ReplaceAllStringFunc(input, func(s string) string {
		return string(s[0] - 'a' + 'A')
	})
	return string(camelCaseStr[0]-'A'+'a') + camelCaseStr[1:]
}
