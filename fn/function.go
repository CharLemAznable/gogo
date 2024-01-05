package fn

type Function[T any, R any] interface {
	Apply(T) R
	CheckedApply(T) (R, error)
}

type FunctionFn[T any, R any] func(T) R

func (fn FunctionFn[T, R]) Apply(t T) R {
	return fn(t)
}

func (fn FunctionFn[T, R]) CheckedApply(t T) (R, error) {
	return fn.Apply(t), nil
}

type FunctionCheckedFn[T any, R any] func(T) (R, error)

func (fn FunctionCheckedFn[T, R]) Apply(t T) R {
	r, _ := fn.CheckedApply(t)
	return r
}

func (fn FunctionCheckedFn[T, R]) CheckedApply(t T) (R, error) {
	return fn(t)
}

func FunctionOf[T any, R any](fn func(T) R) Function[T, R] {
	return FunctionFn[T, R](fn)
}

func FunctionCast[T any, R any](fn func(T) (R, error)) Function[T, R] {
	return FunctionCheckedFn[T, R](fn)
}

func Identity[T any]() Function[T, T] {
	return FunctionOf(func(t T) T { return t })
}

func YCombinator[T any](f func(func(T) T) func(T) T) Function[T, T] {
	return FunctionOf(func(t T) T { return f(YCombinator(f).Apply)(t) })
}
