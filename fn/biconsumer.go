package fn

type BiConsumer[T any, U any] interface {
	Accept(T, U)
	CheckedAccept(T, U) error
}

type BiConsumerFn[T any, U any] func(T, U)

func (fn BiConsumerFn[T, U]) Accept(t T, u U) {
	fn(t, u)
}

func (fn BiConsumerFn[T, U]) CheckedAccept(t T, u U) error {
	fn.Accept(t, u)
	return nil
}

type BiConsumerCheckedFn[T any, U any] func(T, U) error

func (fn BiConsumerCheckedFn[T, U]) Accept(t T, u U) {
	_ = fn.CheckedAccept(t, u)
}

func (fn BiConsumerCheckedFn[T, U]) CheckedAccept(t T, u U) error {
	return fn(t, u)
}

func BiConsumerOf[T any, U any](fn func(T, U)) BiConsumer[T, U] {
	return BiConsumerFn[T, U](fn)
}

func BiConsumerCast[T any, U any](fn func(T, U) error) BiConsumer[T, U] {
	return BiConsumerCheckedFn[T, U](fn)
}
