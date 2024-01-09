package ext_test

import (
	"github.com/CharLemAznable/gogo/ext"
	"github.com/CharLemAznable/gogo/fn"
	"github.com/CharLemAznable/gogo/lang"
	"testing"
)

func TestSliceWithItemFunc(t *testing.T) {
	s := []string{"***", "**", "*"}
	ret := ext.SliceWithItemFunc(s,
		fn.FunctionOf(func(t string) int { return len(t) }))
	if !lang.Equal(ret, []int{3, 2, 1}) {
		t.Errorf("Expected ret is []int{3, 2, 1}, but got '%v'", ret)
	}
}

func TestMapWithItemKeyFunc(t *testing.T) {
	s := []string{"***", "**", "*"}
	ret := ext.MapWithItemKeyFunc(s,
		fn.FunctionOf(func(t string) int { return len(t) }))
	if !lang.Equal(ret, map[int]string{3: "***", 2: "**", 1: "*"}) {
		t.Errorf("Expected ret is map[int]string{3: '***', 2: '**', 1: '*'}, but got '%v'", ret)
	}
}

func TestMapWithItemKeyValueFunc(t *testing.T) {
	s := []string{"***", "**", "*"}
	ret := ext.MapWithItemKeyValueFunc(s,
		fn.FunctionOf(func(t string) int { return len(t) }),
		fn.FunctionOf(func(t string) string { return t + t }))
	if !lang.Equal(ret, map[int]string{3: "******", 2: "****", 1: "**"}) {
		t.Errorf("Expected ret is map[int]string{3: '******', 2: '****', 1: '**'}, but got '%v'", ret)
	}
}
