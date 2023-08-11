package validx

import (
	"strings"

	caseconvert "github.com/Natchalit/gin-x-v1/case-convert"
)

func IsContains(s *[]string, c *string) bool {

	// match case
	for _, v := range *s {
		if v == *c {
			return true
		}
	}
	// lower
	cLower := strings.ToLower(*c)
	if cLower != *c {
		for _, v := range *s {
			if v == cLower {
				return true
			}
		}
	}
	// upper
	cUpper := strings.ToUpper(*c)
	if cUpper != *c {
		for _, v := range *s {
			if v == cUpper {
				return true
			}
		}
	}
	// snake-case
	cSnake := caseconvert.ToSnake(*c)
	if cSnake != *c && cSnake != cLower {
		for _, v := range *s {
			if v == cSnake {
				return true
			}
		}
	}

	return false
}
