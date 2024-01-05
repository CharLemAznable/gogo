package ext_test

import (
	"github.com/CharLemAznable/gogo/ext"
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

	v2 := map[string]string{"a": "A", "b": "B"}
	def2 := map[string]string{"b": "BB", "c": "C"}
	ret2 := ext.MapWithDefault(v2, def2)
	if !lang.Equal(ret2["a"], "A") {
		t.Errorf("Expected ret2['a'] is 'A', but got '%s'", ret1["a"])
	}
	if !lang.Equal(ret2["b"], "B") {
		t.Errorf("Expected ret2['b'] is 'B', but got '%s'", ret1["b"])
	}
	if !lang.Equal(ret2["c"], "C") {
		t.Errorf("Expected ret2['c'] is 'C', but got '%s'", ret1["c"])
	}
}
