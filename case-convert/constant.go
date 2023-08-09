package caseconvert

import (
	"regexp"
	"strings"
)

func ToConstant(input string) string {
	// Convert to CONSTANT_CASE using regex
	re := regexp.MustCompile(`[a-zA-Z]`)
	input = strings.ReplaceAll(input, ` `, ``)
	constantCaseStr := re.ReplaceAllStringFunc(input, func(s string) string {
		return strings.ToUpper(s)
	})
	return constantCaseStr
}

func PascalToConstant(input string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	constantCaseStr := re.ReplaceAllString(input, "${1}_${2}")
	return strings.ToUpper(constantCaseStr)
}

func CamelToConstant(input string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	constantCaseStr := re.ReplaceAllString(input, "${1}_${2}")
	return strings.ToUpper(constantCaseStr)
}

func SnakeToConstant(input string) string {
	return strings.ToUpper(input)
}
