package fn_test

import (
	. "github.com/CharLemAznable/gogo/fn"
	"testing"
)

func TestCheckZeroAccept(t *testing.T) {
	called1 := false
	fn1 := ConsumerOf(func(t string) {
		called1 = true
	})
	called2 := false
	fn2 := RunnableOf(func() {
		called2 = true
	})

	notZeroThenAccept := NotZeroThenAccept(fn1)
	zeroThenRun := IsZeroThenRun[string](fn2)

	var s string

	notZeroThenAccept.Accept(s)
	if called1 {
		t.Error("Expected called1 false, but not")
	}
	s = "abc"
	notZeroThenAccept.Accept(s)
	if !called1 {
		t.Error("Expected called1 true, but not")
	}

	zeroThenRun.Accept(s)
	if called2 {
		t.Error("Expected called2 false, but not")
	}
	s = ""
	zeroThenRun.Accept(s)
	if !called2 {
		t.Error("Expected called2 true, but not")
	}
}

func TestCheckZeroApply(t *testing.T) {
	fn1 := FunctionOf(func(t string) string {
		return "[" + t + "]"
	})
	fn2 := SupplierOf(func() string {
		return "def"
	})

	notZeroThenApply := NotZeroThenApply(fn1)
	zeroThenGet := IsZeroThenGet(fn2)

	var s string

	if "" != notZeroThenApply.Apply(s) {
		t.Error("Expected return '', but not")
	}
	s = "abc"
	if "[abc]" != notZeroThenApply.Apply(s) {
		t.Error("Expected return '[abc]', but not")
	}

	if "abc" != zeroThenGet.Apply(s) {
		t.Error("Expected return 'abc', but not")
	}
	s = ""
	if "def" != zeroThenGet.Apply(s) {
		t.Error("Expected return 'def', but not")
	}
}
