package ext

import "sort"

type Ordered interface {
	Order() string
}

type OrderedSlice []Ordered

func (x OrderedSlice) Len() int           { return len(x) }
func (x OrderedSlice) Less(i, j int) bool { return x[i].Order() < x[j].Order() }
func (x OrderedSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x OrderedSlice) Sort()              { sort.Sort(x) }

func JoinOrdered(items ...Ordered) OrderedSlice {
	return items
}
