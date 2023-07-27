package tox

import (
	"reflect"
	"strings"
)

var TRUE = true
var FALSE = false

func Bool(b interface{}) bool {
	pc := ProcessCovertBoolPtr(b)
	if pc == nil {
		return FALSE
	}

	return *pc
}

func BoolPtr(b interface{}) *bool {
	pc := ProcessCovertBoolPtr(b)
	if pc == nil {
		return &FALSE
	}

	return pc
}

func ProcessCovertBoolPtr(s interface{}) *bool {
	// nil ?
	if s == nil {
		return nil
	}
	// pointer
	if reflect.TypeOf(s).Kind() == reflect.Ptr {
		v := reflect.ValueOf(s)
		if v.IsNil() {
			return nil
		}
		s = reflect.Indirect(v).Interface()
	}
	// bool
	if v, ok := s.(bool); ok {
		return &v
	}
	// int
	if v, ok := s.(int); ok {
		if v == 1 {
			return &TRUE
		}
		return &FALSE
	}
	// string
	if v, ok := s.(string); ok {
		switch strings.TrimSpace(strings.ToUpper(v)) {
		case `1`, `T`, `TRUE`, `Y`, `YES`:
			return &TRUE
		default:
			return &FALSE
		}
	}
	return nil
}
