package validx

import "regexp"

func IsPascalCase(s string) bool {
	re := regexp.MustCompile("^[A-Z][a-z0-9]*([A-Z][a-z0-9]*)*$")
	return re.MatchString(s)
}

func IsSnakeCase(s string) bool {
	re := regexp.MustCompile("^[a-z]+(_[a-z]+)*$")
	return re.MatchString(s)
}

func IsCamelCase(s string) bool {
	re := regexp.MustCompile("^[a-z]+([A-Z][a-z]*)*$")
	return re.MatchString(s)
}
