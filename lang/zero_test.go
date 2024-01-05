package lang_test

import (
	"github.com/CharLemAznable/gogo/lang"
	"testing"
)

type myFunc func()

func TestIsZero(t *testing.T) {
	if !lang.IsZero(nil) {
		t.Error("Expected isZero(nil), but not")
	}

	var i int
	if !lang.IsZero(i) {
		t.Error("Expected isZero(int(nil)), but not")
	}
	if lang.IsZero(1) {
		t.Error("Expected not isZero(1), but not")
	}

	var s string
	if !lang.IsZero(s) {
		t.Error("Expected isZero(string(nil)), but not")
	}
	if !lang.IsZero("") {
		t.Error("Expected isZero(''), but not")
	}
	if lang.IsZero("nil") {
		t.Error("Expected not isZero('nil'), but not")
	}

	var f myFunc
	if !lang.IsZero(f) {
		t.Error("Expected isZero(myFunc(nil)), but not")
	}
}
