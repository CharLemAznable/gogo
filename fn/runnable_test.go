package fn_test

import (
	"errors"
	. "github.com/CharLemAznable/gogo/fn"
	"testing"
)

func TestRunnableOf(t *testing.T) {
	called := false
	fn := func() {
		called = true
	}

	r := RunnableOf(fn)
	err := r.CheckedRun()
	if err != nil {
		t.Errorf("Expected error is nil, but got '%v'", err)
	}
	if !called {
		t.Error("Expected called, but not called")
	}
}

func TestRunnableCast(t *testing.T) {
	called := false
	err := errors.New("error")
	fn := func() error {
		called = true
		return err
	}

	r := RunnableCast(fn)
	r.Run()
	if !called {
		t.Error("Expected called, but not called")
	}
}
