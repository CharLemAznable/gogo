package lang_test

import (
	"github.com/CharLemAznable/gogo/lang"
	"testing"
)

func TestEquals(t *testing.T) {
	if !lang.Equal(nil, nil) {
		t.Error("Expected equals(nil, nil), but not")
	}
	if lang.Equal(1, nil) {
		t.Error("Expected not equals(1, nil), but not")
	}
	if lang.Equal(nil, "") {
		t.Error("Expected not equals(nil, ''), but not")
	}
	if lang.Equal(1, "1") {
		t.Error("Expected not equals(1, '1'), but not")
	}
	if !lang.Equal(1, 1) {
		t.Error("Expected equals(1, 1), but not")
	}
	if !lang.Equal("1", "1") {
		t.Error("Expected equals('1', '1'), but not")
	}

	fn1 := func() {}
	fn2 := func() {}
	if lang.Equal(fn1, fn2) {
		t.Error("Expected not equals(fn1, fn2), but not")
	}
	if !lang.Equal(fn1, fn1) {
		t.Error("Expected equals(fn1, fn1), but not")
	}
}
