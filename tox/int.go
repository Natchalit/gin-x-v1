package tox

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

var INT int
var INT32 int32
var INT64 int64

func Int(n interface{}) int {
	pc := ProcessConvertIntPtr(n)
	if pc == nil {
		return INT
	}

	return int(*pc)
}

func IntPtr(n interface{}) *int {
	pc := ProcessConvertIntPtr(n)
	if pc == nil {
		return &INT
	}

	pcx := int(*pc)

	return &pcx
}

func Int32(n interface{}) int32 {
	pc := ProcessConvertIntPtr(n)
	if pc == nil {
		return INT32
	}
	return int32(*pc)
}

func Int64(n interface{}) int64 {
	pc := ProcessConvertIntPtr(n)
	if pc == nil {
		return INT64
	}

	return *pc
}

func IntPtr64(n interface{}) *int64 {
	pc := ProcessConvertIntPtr(n)
	if pc == nil {
		return &INT64
	}

	return pc
}

func ProcessConvertIntPtr(n interface{}) *int64 {

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
		r := int64(n.(int))
		return &r
	case reflect.Int8:
		r := int64(n.(int8))
		return &r
	case reflect.Int16:
		r := int64(n.(int16))
		return &r
	case reflect.Int32:
		r := int64(n.(int32))
		return &r
	case reflect.Int64:
		r := n.(int64)
		return &r
	case reflect.Float32:
		r := int64(math.Round(n.(float64)))
		return &r
	case reflect.Float64:
		r := int64(math.Round(n.(float64)))
		return &r
	case reflect.Uint:
		r := int64(n.(uint))
		return &r
	case reflect.Uint8:
		r := int64(n.(uint8))
		return &r
	case reflect.Uint16:
		r := int64(n.(uint16))
		return &r
	case reflect.Uint32:
		r := int64(n.(uint32))
		return &r
	case reflect.Uint64:
		r := int64(n.(uint64))
		return &r
	}
	nx, ex := strconv.ParseInt(fmt.Sprintf(`%v`, n.(string)), 10, 64)
	if ex != nil {
		return &INT64
	}
	r := int64(nx)

	return &r
}
