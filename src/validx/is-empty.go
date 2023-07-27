package validx

import "strings"

func IsEmpty(s string) bool {
	s = strings.TrimSpace(s)
	return IsEmptyPtr(&s)
}

func IsEmptyPtr(s *string) bool {
	if s == nil || *s == `` {
		return true
	}
	return false
}
