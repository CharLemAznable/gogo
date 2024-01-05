package fn_test

import (
	. "github.com/CharLemAznable/gogo/fn"
	"testing"
)

func TestBiConsumerOf(t *testing.T) {
	fn := func(i int, s string) {
		// do something
	}

	con := BiConsumerOf(fn)
	err := con.CheckedAccept(10, "test")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestBiConsumerCast(t *testing.T) {
	called := false
	fn := func(i int, s string) error {
		called = true
		return nil
	}

	con := BiConsumerCast(fn)
	con.Accept(10, "test")
	if !called {
		t.Error("Expected called, but not called")
	}
}
