package fn

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

func ConsumerOf[T any](fn func(T)) Consumer[T] {
	return ConsumerFn[T](fn)
}

func ConsumerCast[T any](fn func(T) error) Consumer[T] {
	return ConsumerCheckedFn[T](fn)
}

func ConsumerChan[T any](ch chan T) Consumer[T] {
	return ConsumerCh[T](ch)
}
