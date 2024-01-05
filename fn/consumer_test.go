package fn_test

import (
	"errors"
	. "github.com/CharLemAznable/gogo/fn"
	"sync"
	"sync/atomic"
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

func TestConsumers(t *testing.T) {
	wg := &sync.WaitGroup{}
	var sum atomic.Int64

	wg.Add(1)
	cFn := ConsumerOf(func(i int) {
		sum.Add(int64(i))
		wg.Done()
	})

	wg.Add(1)
	ch := make(chan int, 1)
	cCh := ConsumerChan(ch)
	go func() {
		select {
		case i := <-ch:
			sum.Add(int64(i))
			wg.Done()
		}
	}()

	JoinConsumers(cFn, cCh).Accept(2)
	wg.Wait()
	if sum.Load() != 4 {
		t.Errorf("Expected sum 4, but got %d", sum.Load())
	}
}

func TestConsumersChecked(t *testing.T) {
	wg := &sync.WaitGroup{}
	var sum atomic.Int64

	wg.Add(1)
	ch1 := make(chan int, 1)
	c1 := ConsumerChan(ch1)
	go func() {
		select {
		case i := <-ch1:
			sum.Add(int64(i))
			wg.Done()
		}
	}()

	wg.Add(1)
	ch2 := make(chan int, 1)
	c2 := ConsumerChan(ch2)
	go func() {
		select {
		case i := <-ch2:
			sum.Add(int64(i))
			wg.Done()
		}
	}()

	err := JoinConsumers(c1, c2).CheckedAccept(2)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	wg.Wait()
	if sum.Load() != 4 {
		t.Errorf("Expected sum 4, but got %d", sum.Load())
	}
}
