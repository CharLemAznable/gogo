package fn

import (
	"github.com/CharLemAznable/gogo/lang"
	"sync"
)

type Consumer[T any] interface {
	Accept(T)
	CheckedAccept(T) error
}

type ConsumerFn[T any] func(T)

func (fn ConsumerFn[T]) Accept(t T) {
	fn(t)
}

func (fn ConsumerFn[T]) CheckedAccept(t T) error {
	fn.Accept(t)
	return nil
}

type ConsumerCheckedFn[T any] func(T) error

func (fn ConsumerCheckedFn[T]) Accept(t T) {
	_ = fn.CheckedAccept(t)
}

func (fn ConsumerCheckedFn[T]) CheckedAccept(t T) error {
	return fn(t)
}

type ConsumerCh[T any] chan<- T

func (ch ConsumerCh[T]) Accept(t T) {
	ch <- t
}

func (ch ConsumerCh[T]) CheckedAccept(t T) error {
	ch.Accept(t)
	return nil
}

type Consumers[T any] []Consumer[T]

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
		go func(c Consumer[T]) {
			err.Append(c.CheckedAccept(t))
			wg.Done()
		}(sub)
	}
	wg.Wait()
	return err.MaybeUnwrap()
}

func ConsumerOf[T any](fn func(T)) Consumer[T] {
	return ConsumerFn[T](fn)
}

func ConsumerCast[T any](fn func(T) error) Consumer[T] {
	return ConsumerCheckedFn[T](fn)
}

func ConsumerChan[T any](ch chan T) Consumer[T] {
	return ConsumerCh[T](ch)
}

func Ignore[T any]() Consumer[T] {
	return ConsumerOf(func(t T) { /* ignore */ })
}

func JoinConsumers[T any](consumers ...Consumer[T]) Consumer[T] {
	return Consumers[T](consumers)
}
