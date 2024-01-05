package ext

import . "github.com/CharLemAznable/gogo/fn"

func CheckZeroRun[T any](value T, zeroFn func(), notZeroFn func(T)) {
	CheckZeroAccept(RunnableOf(zeroFn), ConsumerOf(notZeroFn)).Accept(value)
}

func NotZeroThenRun[T any](value T, notZeroFn func(T)) {
	CheckZeroRun(value, func() {}, notZeroFn)
}

func ZeroThenRun[T any](value T, zeroFn func()) {
	CheckZeroRun(value, zeroFn, func(t T) {})
}

func CheckZero[T any, R any](value T, zeroFn func() R, notZeroFn func(T) R) R {
	return CheckZeroApply(SupplierOf(zeroFn), FunctionOf(notZeroFn)).Apply(value)
}

func NotZeroThen[T any, R any](value T, notZeroFn func(T) R) R {
	return CheckZero(value, Zero[R]().Get, notZeroFn)
}

func ZeroThen[T any](value T, zeroFn func() T) T {
	return CheckZero(value, zeroFn, Identity[T]().Apply)
}

func CheckEmptyRun(value string, emptyFn func(), notEmptyFn func(string)) {
	CheckZeroRun(value, emptyFn, notEmptyFn)
}

func NotEmptyThenRun(value string, notEmptyFn func(string)) {
	NotZeroThenRun(value, notEmptyFn)
}

func EmptyThenRun(value string, emptyFn func()) {
	ZeroThenRun(value, emptyFn)
}

func CheckEmpty[R any](value string, emptyFn func() R, notEmptyFn func(string) R) R {
	return CheckZero(value, emptyFn, notEmptyFn)
}

func NotEmptyThen[R any](value string, notEmptyFn func(string) R) R {
	return NotZeroThen(value, notEmptyFn)
}

func EmptyThen(value string, emptyFn func() string) string {
	return ZeroThen(value, emptyFn)
}
