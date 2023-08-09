package caseconvert

import (
	"regexp"
	"strings"
)

func ToSnake(input string) string {
	// Convert to snake_case using regex
	re := regexp.MustCompile(`[A-Z]`)
	input = strings.ReplaceAll(input, ` `, ``)
	snakeCaseStr := re.ReplaceAllStringFunc(input, func(s string) string {
		return "_" + strings.ToLower(s)
	})
	return strings.TrimLeft(snakeCaseStr, "_")
}

func PascalToSnake(input string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snakeCaseStr := re.ReplaceAllString(input, "${1}_${2}")
	return snakeCaseStr
}

func CamelToSnake(input string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snakeCaseStr := re.ReplaceAllString(input, "${1}_${2}")
	return snakeCaseStr
}

func ConstantToSnake(input string) string {
	return strings.ToLower(input)
}
