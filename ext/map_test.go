package ext_test

import (
	"github.com/CharLemAznable/gogo/ext"
	"github.com/CharLemAznable/gogo/fn"
	"github.com/CharLemAznable/gogo/lang"
	"testing"
)

func TestMapWithDefault(t *testing.T) {
	f := func() {}
	v1 := map[string]any{"a": 1, "b": f}
	def1 := map[string]any{"b": "B", "c": "C"}
	ret1 := ext.MapWithDefault(v1, def1)
	if !lang.Equal(ret1["a"], 1) {
		t.Errorf("Expected ret1['a'] is 1, but got '%v'", ret1["a"])
	}
	if !lang.Equal(ret1["b"], f) {
		t.Errorf("Expected ret1['b'] is func() {}, but got '%v'", ret1["b"])
	}
	if !lang.Equal(ret1["c"], "C") {
		t.Errorf("Expected ret1['c'] is 'C', but got '%v'", ret1["c"])
	}

	v2 := map[int]string{1: "A", 2: "B"}
	def2 := map[int]string{2: "BB", 3: "C"}
	ret2 := ext.MapWithDefault(v2, def2)
	if !lang.Equal(ret2[1], "A") {
		t.Errorf("Expected ret2[1] is 'A', but got '%s'", ret2[1])
	}
	if !lang.Equal(ret2[2], "B") {
		t.Errorf("Expected ret2[2] is 'B', but got '%s'", ret2[2])
	}
	if !lang.Equal(ret2[3], "C") {
		t.Errorf("Expected ret2[3] is 'C', but got '%s'", ret2[3])
	}
}

func TestMapWithValueFunc(t *testing.T) {
	v := map[string]string{"a": "*", "b": "**", "c": "***"}
	ret := ext.MapWithValueFunc(v, fn.FunctionOf(func(val string) int { return len(val) }))
	if !lang.Equal(ret["a"], 1) {
		t.Errorf("Expected ret['a'] is 1, but got '%d'", ret["a"])
	}
	if !lang.Equal(ret["b"], 2) {
		t.Errorf("Expected ret['b'] is 2, but got '%d'", ret["b"])
	}
	if !lang.Equal(ret["c"], 3) {
		t.Errorf("Expected ret['c'] is 3, but got '%d'", ret["c"])
	}
}

func TestMapWithKeyValueFunc(t *testing.T) {
	v := map[int]string{1: "*", 2: "**", 3: "***"}
	ret := ext.MapWithKeyValueFunc(v, fn.BiFunctionOf(func(key int, val string) int { return len(val) }))
	if !lang.Equal(ret[1], 1) {
		t.Errorf("Expected ret[1] is 1, but got '%d'", ret[1])
	}
	if !lang.Equal(ret[2], 2) {
		t.Errorf("Expected ret[2] is 2, but got '%d'", ret[2])
	}
	if !lang.Equal(ret[3], 3) {
		t.Errorf("Expected ret[3] is 3, but got '%d'", ret[3])
	}
}
