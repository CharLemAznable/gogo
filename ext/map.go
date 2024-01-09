package ext

import "github.com/CharLemAznable/gogo/fn"

func MapWithDefault[K string | int, V any](v, def map[K]V) map[K]V {
	ret := make(map[K]V)
	for key, value := range def {
		ret[key] = value
	}
	for key, value := range v {
		ret[key] = value
	}
	return ret
}

func MapWithValueFunc[K string | int, V any, R any](
	v map[K]V, mapper fn.Function[V, R]) map[K]R {
	ret := make(map[K]R)
	for key, value := range v {
		ret[key] = mapper.Apply(value)
	}
	return ret
}

func MapWithKeyValueFunc[K string | int, V any, R any](
	v map[K]V, mapper fn.BiFunction[K, V, R]) map[K]R {
	ret := make(map[K]R)
	for key, value := range v {
		ret[key] = mapper.Apply(key, value)
	}
	return ret
}
