package fn_test

import (
	"errors"
	. "github.com/CharLemAznable/gogo/fn"
	"testing"
)

func TestConsumerOf(t *testing.T) {
	called := false
	fn := func(i int) {
		called = true
	}

	con := ConsumerOf(fn)
	err := con.CheckedAccept(10)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if !called {
		t.Error("Expected called, but not called")
	}
}

func TestConsumerCast(t *testing.T) {
	called := false
	fn := func(i int) error {
		called = true
		return errors.New("error")
	}

	con := ConsumerCast(fn)
	con.Accept(10)
	if !called {
		t.Error("Expected called, but not called")
	}
}

func TestConsumerChan(t *testing.T) {
	ch := make(chan int, 1)
	cn := ConsumerChan(ch)
	err := cn.CheckedAccept(10)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	select {
	case i := <-ch:
		if i != 10 {
			t.Errorf("Expected get 10, but got %d", i)
		}
	}
}
