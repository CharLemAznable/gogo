package fn

import "github.com/CharLemAznable/gogo/lang"

type Supplier[T any] interface {
	Get() T
	CheckedGet() (T, error)
}

type SupplierFn[T any] func() T

func (fn SupplierFn[T]) Get() T {
	return fn()
}

func (fn SupplierFn[T]) CheckedGet() (T, error) {
	return fn.Get(), nil
}

type SupplierCheckedFn[T any] func() (T, error)

func (fn SupplierCheckedFn[T]) Get() T {
	t, _ := fn.CheckedGet()
	return t
}

func (fn SupplierCheckedFn[T]) CheckedGet() (T, error) {
	return fn()
}

func SupplierOf[T any](fn func() T) Supplier[T] {
	return SupplierFn[T](fn)
}

func SupplierCast[T any](fn func() (T, error)) Supplier[T] {
	return SupplierCheckedFn[T](fn)
}

func Constant[T any](t T) Supplier[T] {
	return SupplierOf(func() T { return t })
}

func Zero[T any]() Supplier[T] {
	return Constant(lang.Zero[T]())
}
