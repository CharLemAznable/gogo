package fn

type BiFunction[T any, U any, R any] interface {
	Apply(T, U) R
	CheckedApply(T, U) (R, error)
	Curry() func(T) Function[U, R]
	Partial(T) Function[U, R]
}

type BiFunctionFn[T any, U any, R any] func(T, U) R

func (fn BiFunctionFn[T, U, R]) Apply(t T, u U) R {
	return fn(t, u)
}

func (fn BiFunctionFn[T, U, R]) CheckedApply(t T, u U) (R, error) {
	return fn.Apply(t, u), nil
}

func (fn BiFunctionFn[T, U, R]) Curry() func(T) Function[U, R] {
	return func(t T) Function[U, R] {
		return FunctionOf(func(u U) R { return fn(t, u) })
	}
}

func (fn BiFunctionFn[T, U, R]) Partial(t T) Function[U, R] {
	return FunctionOf(func(u U) R { return fn(t, u) })
}

type BiFunctionCheckedFn[T any, U any, R any] func(T, U) (R, error)

func (fn BiFunctionCheckedFn[T, U, R]) Apply(t T, u U) R {
	r, _ := fn.CheckedApply(t, u)
	return r
}

func (fn BiFunctionCheckedFn[T, U, R]) CheckedApply(t T, u U) (R, error) {
	return fn(t, u)
}

func (fn BiFunctionCheckedFn[T, U, R]) Curry() func(T) Function[U, R] {
	return func(t T) Function[U, R] {
		return FunctionCast(func(u U) (R, error) { return fn(t, u) })
	}
}

func (fn BiFunctionCheckedFn[T, U, R]) Partial(t T) Function[U, R] {
	return FunctionCast(func(u U) (R, error) { return fn(t, u) })
}

func BiFunctionOf[T any, U any, R any](fn func(T, U) R) BiFunction[T, U, R] {
	return BiFunctionFn[T, U, R](fn)
}

func BiFunctionCast[T any, U any, R any](fn func(T, U) (R, error)) BiFunction[T, U, R] {
	return BiFunctionCheckedFn[T, U, R](fn)
}
