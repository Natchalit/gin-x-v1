package caseconvert

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnake(input string) string {
	snake := matchFirstCap.ReplaceAllString(input, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
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
