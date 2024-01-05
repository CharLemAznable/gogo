package lang_test

import (
	"github.com/CharLemAznable/gogo/lang"
	"testing"
)

type testFn func()

func TestRemoveElementByValue_Func(t *testing.T) {
	fn1 := func() {}
	fn2 := func() {}
	slice := []testFn{fn1, fn2, fn2, fn1}
	result := lang.RemoveElementByValue(slice, fn2)
	expected := []testFn{fn1, fn1}
	if !lang.Equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
	fn3 := func() {}
	result = lang.RemoveElementByValue(slice, fn3)
	expected = []testFn{fn1, fn2, fn2, fn1}
	if !lang.Equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
	result = lang.AppendElementUnique(slice, fn2)
	expected = []testFn{fn1, fn1, fn2}
	if !lang.Equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
