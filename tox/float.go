package tox

import (
	"fmt"
	"reflect"
	"strconv"
)

var FLOAT32 float32
var FLOAT64 float64

func Float32(n interface{}) float32 {
	r := ProcessConvertFloatPtr(n)
	if r == nil {
		return FLOAT32
	}
	return float32(*r)
}
func Float64(n interface{}) float64 {
	r := ProcessConvertFloatPtr(n)
	if r == nil {
		return FLOAT64
	}
	return *r
}

func Float32Ptr(n interface{}) *float32 {
	r := ProcessConvertFloatPtr(n)
	if r == nil {
		return &FLOAT32
	}
	rf := float32(*r)
	return &rf
}
func Float64Ptr(n interface{}) *float64 {
	r := ProcessConvertFloatPtr(n)
	if r == nil {
		return &FLOAT64
	}
	return r
}

func ProcessConvertFloatPtr(n interface{}) *float64 {

	if n == nil {
		return nil
	}

	// Is pointer?
	if reflect.TypeOf(n).Kind() == reflect.Pointer {
		vn := reflect.ValueOf(n)
		// Has value?
		if vn.IsNil() {
			return nil
		}
		n = reflect.Indirect(vn).Interface()
	}

	switch reflect.TypeOf(n).Kind() {
	case reflect.Int:
		r := float64(n.(int))
		return &r
	case reflect.Int8:
		r := float64(n.(int8))
		return &r
	case reflect.Int16:
		r := float64(n.(int16))
		return &r
	case reflect.Int32:
		r := float64(n.(int32))
		return &r
	case reflect.Int64:
		r := float64(n.(int64))
		return &r
	case reflect.Float32:
		r := Float64(n.(float64))
		return &r
	case reflect.Float64:
		r := n.(float64)
		return &r
	case reflect.Uint:
		r := float64(n.(uint))
		return &r
	case reflect.Uint8:
		r := float64(n.(uint8))
		return &r
	case reflect.Uint16:
		r := float64(n.(uint16))
		return &r
	case reflect.Uint32:
		r := float64(n.(uint32))
		return &r
	case reflect.Uint64:
		r := float64(n.(uint64))
		return &r
	}

	nx, ex := strconv.ParseFloat(fmt.Sprintf(`%v`, n.(string)), 64)
	if ex != nil {
		return &FLOAT64
	}
	r := float64(nx)
	return &r
}
