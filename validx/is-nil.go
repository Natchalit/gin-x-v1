package validx

import "reflect"

func IsNil(v any) bool {

	// https://mangatmodi.medium.com/go-check-nil-interface-the-right-way-d142776edef1

	if v == nil {
		return true
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(v).IsNil()
	}
	return false

}
