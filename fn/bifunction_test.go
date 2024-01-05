package fn_test

import (
	. "github.com/CharLemAznable/gogo/fn"
	"testing"
)

func TestBiFunctionOf(t *testing.T) {
	fn := BiFunctionOf(func(a int, b string) int {
		return a + len(b)
	})

	result, err := fn.CheckedApply(5, "hello")
	expected := 10
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestBiFunctionCast(t *testing.T) {
	fn := BiFunctionCast(func(a int, b string) (int, error) {
		return a + len(b), nil
	})

	result := fn.Apply(5, "hello")
	expected := 10
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestBiFunctionFnCurry(t *testing.T) {
	fn := BiFunctionOf(func(t int, u string) bool {
		return true
	})

	curriedFn := fn.Curry()(1)
	r := curriedFn.Apply("test")
	if r != true {
		t.Errorf("TestCurry failed")
	}
}

func TestBiFunctionFnPartial(t *testing.T) {
	fn := BiFunctionOf(func(t int, u string) bool {
		return true
	})

	partialFn := fn.Partial(1)
	r := partialFn.Apply("test")
	if r != true {
		t.Errorf("TestPartial failed")
	}
}

func TestBiFunctionCheckedFnCurry(t *testing.T) {
	fn := BiFunctionCast(func(t int, u string) (bool, error) {
		return true, nil
	})

	curriedFn := fn.Curry()(1)
	r, err := curriedFn.CheckedApply("test")
	if r != true || err != nil {
		t.Errorf("TestCurry failed")
	}
}

func TestBiFunctionCheckedFnPartial(t *testing.T) {
	fn := BiFunctionCast(func(t int, u string) (bool, error) {
		return true, nil
	})

	partialFn := fn.Partial(1)
	r, err := partialFn.CheckedApply("test")
	if r != true || err != nil {
		t.Errorf("TestPartial failed")
	}
}
