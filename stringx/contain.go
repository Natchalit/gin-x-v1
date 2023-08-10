package stringx

import caseconvert "github.com/Natchalit/gin-x-v1/case-convert"

func IsContain(arr []string, target string) bool {
	for _, element := range arr {
		if caseconvert.ToSnake(element) == target {
			return true
		}
	}
	return false
}
