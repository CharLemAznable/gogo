package lang

import "reflect"

func IsZero(val any) bool {
	if val == nil {
		return true
	}
	return Equal(val, reflect.Zero(
		reflect.TypeOf(val)).Interface())
}
