package ext

func MapWithDefault[V any](v, def map[string]V) map[string]V {
	ret := make(map[string]V)
	for key, value := range def {
		ret[key] = value
	}
	for key, value := range v {
		ret[key] = value
	}
	return ret
}
