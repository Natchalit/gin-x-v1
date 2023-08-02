package caseconvert

import (
	"regexp"
	"strings"
)

func ToPascal(input string) string {
	// Convert to PascalCase using regex
	re := regexp.MustCompile(`(?:^|_)([a-zA-Z])`)
	pascalCaseStr := re.ReplaceAllStringFunc(input, func(s string) string {
		return strings.ToUpper(string(s[1]))
	})
	return pascalCaseStr
}

func CamelToPascal(input string) string {
	re := regexp.MustCompile(`^[a-z]`)
	pascalCaseStr := re.ReplaceAllStringFunc(input, func(s string) string {
		return string(s[0] - 'a' + 'A')
	})
	return pascalCaseStr
}

func SnakeToPascal(input string) string {
	re := regexp.MustCompile(`_([a-z0-9])`)
	pascalCaseStr := re.ReplaceAllStringFunc(input, func(s string) string {
		return string(s[1] - 'a' + 'A')
	})
	return string(pascalCaseStr[0]-'a'+'A') + pascalCaseStr[1:]
}

func ConstantToPascal(input string) string {
	re := regexp.MustCompile(`_([a-zA-Z0-9])`)
	pascalCaseStr := re.ReplaceAllStringFunc(input, func(s string) string {
		return string(s[1] - 'a' + 'A')
	})
	return pascalCaseStr
}
