package ext_test

import (
	"github.com/CharLemAznable/gogo/ext"
	"github.com/CharLemAznable/gogo/fn"
	"sync/atomic"
	"testing"
	"time"
)

func TestConsumers(t *testing.T) {
	count := atomic.Int64{}
	countFn := func(t string) {
		count.Add(1)
	}
	consumers := ext.NewConsumers[string]()
	consumers.AppendConsumer(fn.ConsumerOf(countFn))
	consumers.Accept("ABC")
	time.Sleep(time.Second)
	if 1 != count.Load() {
		t.Errorf("Expected count 1, but got '%d'", count.Load())
	}
	consumers.RemoveConsumer(fn.ConsumerOf(countFn))
	err := consumers.CheckedAccept("ABC")
	if err != nil {
		t.Errorf("Unexpected err: %s", err.Error())
	}
	time.Sleep(time.Second)
	if 1 != count.Load() {
		t.Errorf("Expected count 1, but got '%d'", count.Load())
	}
}
