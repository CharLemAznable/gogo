package lang

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

type PanicError struct {
	msg string
	v   any
}

func (e *PanicError) Error() string {
	return e.msg
}

func (e *PanicError) Origin() any {
	return e.v
}

func WrapPanic(v any) *PanicError {
	return &PanicError{msg: fmt.Sprintf("panicked with %v", v), v: v}
}

type Panicked chan any

func (pr Panicked) Recover() {
	if err := recover(); err != nil {
		pr <- err
	}
}

func (pr Panicked) Caught() <-chan any {
	return pr
}
