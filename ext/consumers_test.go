package ext_test

import (
	"github.com/CharLemAznable/gogo/ext"
	"github.com/CharLemAznable/gogo/fn"
	"sync"
	"sync/atomic"
	"testing"
)

func TestConsumers(t *testing.T) {
	wg := &sync.WaitGroup{}
	var sum atomic.Int64

	wg.Add(1)
	cFn := fn.ConsumerOf(func(i int) {
		sum.Add(int64(i))
		wg.Done()
	})

	wg.Add(1)
	ch := make(chan int, 1)
	cCh := fn.ConsumerChan(ch)
	go func() {
		select {
		case i := <-ch:
			sum.Add(int64(i))
			wg.Done()
		}
	}()

	ext.JoinConsumers(cFn, cCh).Accept(2)
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
	c1 := fn.ConsumerChan(ch1)
	go func() {
		select {
		case i := <-ch1:
			sum.Add(int64(i))
			wg.Done()
		}
	}()

	wg.Add(1)
	ch2 := make(chan int, 1)
	c2 := fn.ConsumerChan(ch2)
	go func() {
		select {
		case i := <-ch2:
			sum.Add(int64(i))
			wg.Done()
		}
	}()

	err := ext.JoinConsumers(c1, c2).CheckedAccept(2)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	wg.Wait()
	if sum.Load() != 4 {
		t.Errorf("Expected sum 4, but got %d", sum.Load())
	}
}
