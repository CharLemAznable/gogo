package ext

import "sort"

type Ordered interface {
	Order() string
}

type OrderedSlice[T Ordered] []T

func (x OrderedSlice[T]) Len() int           { return len([]T(x)) }
func (x OrderedSlice[T]) Less(i, j int) bool { return []T(x)[i].Order() < []T(x)[j].Order() }
func (x OrderedSlice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x OrderedSlice[T]) Sort()              { sort.Sort(x) }

func JoinOrdered[T Ordered](items ...T) OrderedSlice[T] {
	return items
}
