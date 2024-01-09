package ext

import "github.com/CharLemAznable/gogo/fn"

func SliceWithItemFunc[T any, R any](
	s []T, mapper fn.Function[T, R]) []R {
	ret := make([]R, 0)
	for _, item := range s {
		ret = append(ret, mapper.Apply(item))
	}
	return ret
}

func MapWithItemKeyFunc[T any, K string | int](
	s []T, mapper fn.Function[T, K]) map[K]T {
	ret := make(map[K]T)
	for _, item := range s {
		ret[mapper.Apply(item)] = item
	}
	return ret
}

func MapWithItemKeyValueFunc[T any, K string | int, V any](
	s []T, keyMapper fn.Function[T, K], valueMapper fn.Function[T, V]) map[K]V {
	ret := make(map[K]V)
	for _, item := range s {
		ret[keyMapper.Apply(item)] = valueMapper.Apply(item)
	}
	return ret
}
