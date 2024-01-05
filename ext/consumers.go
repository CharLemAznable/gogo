package ext

import (
	"github.com/CharLemAznable/gogo/fn"
	"github.com/CharLemAznable/gogo/lang"
	"sync"
)

type Consumers[T any] []fn.Consumer[T]

func (s Consumers[T]) Accept(t T) {
	for _, sub := range s {
		go sub.Accept(t)
	}
}

func (s Consumers[T]) CheckedAccept(t T) error {
	wg := &sync.WaitGroup{}
	err := lang.MultiError{}
	for _, sub := range s {
		wg.Add(1)
		go func(c fn.Consumer[T]) {
			err.Append(c.CheckedAccept(t))
			wg.Done()
		}(sub)
	}
	wg.Wait()
	return err.MaybeUnwrap()
}

func JoinConsumers[T any](consumers ...fn.Consumer[T]) fn.Consumer[T] {
	return Consumers[T](consumers)
}
