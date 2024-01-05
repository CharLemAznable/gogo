package lang

import "fmt"

func Cast[T any](val interface{}) (T, error) {
	t, ok := val.(T)
	if !ok {
		return t, fmt.Errorf("%#v type mismatch %T", val, t)
	}
	return t, nil
}

func CastQuietly[T any](val interface{}) T {
	t, _ := val.(T)
	return t
}

func Zero[T any]() T {
	return CastQuietly[T](nil)
}

func CastOrZero[T any](val interface{}) (T, error) {
	if val == nil {
		return Zero[T](), nil
	}
	return Cast[T](val)
}
