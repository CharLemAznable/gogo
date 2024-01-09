package fn

import "github.com/CharLemAznable/gogo/lang"

func MustRun(runnable Runnable) Runnable {
	return RunnableCast(func() error {
		finished := make(chan *retErr)
		panicked := make(lang.Panicked)
		go func() {
			defer panicked.Recover()
			err := runnable.CheckedRun()
			finished <- &retErr{err: err}
		}()
		select {
		case ret := <-finished:
			return ret.err
		case v := <-panicked.Caught():
			return lang.WrapPanic(v)
		}
	})
}

func MustGet[T any](supplier Supplier[T]) Supplier[T] {
	return SupplierCast(func() (T, error) {
		finished := make(chan *retValErr[T])
		panicked := make(lang.Panicked)
		go func() {
			defer panicked.Recover()
			val, err := supplier.CheckedGet()
			finished <- &retValErr[T]{val: val, err: err}
		}()
		select {
		case ret := <-finished:
			return ret.val, ret.err
		case v := <-panicked.Caught():
			return lang.Zero[T](), lang.WrapPanic(v)
		}
	})
}

func MustAccept[T any](consumer Consumer[T]) Consumer[T] {
	return ConsumerCast(func(t T) error {
		finished := make(chan *retErr)
		panicked := make(lang.Panicked)
		go func(t T) {
			defer panicked.Recover()
			err := consumer.CheckedAccept(t)
			finished <- &retErr{err: err}
		}(t)
		select {
		case ret := <-finished:
			return ret.err
		case v := <-panicked.Caught():
			return lang.WrapPanic(v)
		}
	})
}

func MustApply[T any, R any](function Function[T, R]) Function[T, R] {
	return FunctionCast(func(t T) (R, error) {
		finished := make(chan *retValErr[R])
		panicked := make(lang.Panicked)
		go func(t T) {
			defer panicked.Recover()
			val, err := function.CheckedApply(t)
			finished <- &retValErr[R]{val: val, err: err}
		}(t)
		select {
		case ret := <-finished:
			return ret.val, ret.err
		case v := <-panicked.Caught():
			return lang.Zero[R](), lang.WrapPanic(v)
		}
	})
}

type retErr struct {
	err error
}

type retValErr[T any] struct {
	val T
	err error
}
