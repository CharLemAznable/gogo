package fn

import "github.com/CharLemAznable/gogo/lang"

func CheckZeroAccept[T any](
	zeroFn Runnable,
	notZeroFn Consumer[T]) Consumer[T] {
	return ConsumerCast(func(t T) error {
		if lang.IsZero(t) {
			return zeroFn.CheckedRun()
		}
		return notZeroFn.CheckedAccept(t)
	})
}

func NotZeroThenAccept[T any](notZeroFn Consumer[T]) Consumer[T] {
	return CheckZeroAccept(Empty(), notZeroFn)
}

func IsZeroThenRun[T any](zeroFn Runnable) Consumer[T] {
	return CheckZeroAccept(zeroFn, Ignore[T]())
}

func CheckZeroApply[T any, R any](
	zeroFn Supplier[R],
	notZeroFn Function[T, R]) Function[T, R] {
	return FunctionCast(func(t T) (R, error) {
		if lang.IsZero(t) {
			return zeroFn.CheckedGet()
		}
		return notZeroFn.CheckedApply(t)
	})
}

func NotZeroThenApply[T any, R any](
	notZeroFn Function[T, R]) Function[T, R] {
	return CheckZeroApply(Zero[R](), notZeroFn)
}

func IsZeroThenGet[T any](
	zeroFn Supplier[T]) Function[T, T] {
	return CheckZeroApply(zeroFn, Identity[T]())
}
