package lang

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

type PanicError struct {
	v   any
	msg string
}

func (e *PanicError) Origin() any {
	return e.v
}

func (e *PanicError) Error() string {
	return e.msg
}

func ErrorOfPanic(v any) *PanicError {
	return &PanicError{v: v,
		msg: fmt.Sprintf("panicked with %v", v)}
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
