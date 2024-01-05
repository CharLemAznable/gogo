package lang_test

import (
	"errors"
	"github.com/CharLemAznable/gogo/lang"
	"testing"
)

func TestPanicIfError(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Expected no panic, but got %v", r)
			}
		}()
		lang.PanicIfError(nil)
	}()
	func() {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("Expected panic, but got nothing")
			}
			err, ok := r.(error)
			if !ok {
				t.Errorf("Expected panic error, but got %v", r)
			} else if err.Error() != "panic error" {
				t.Errorf("Expected error message 'panic error', but got '%s'", err.Error())
			}
		}()
		lang.PanicIfError(errors.New("panic error"))
	}()
}

func TestPanicRecover(t *testing.T) {
	finished := make(chan error)
	panicked := make(lang.Panicked)

	go func() {
		defer panicked.Recover()
		finished <- errors.New("error")
	}()

	var actualError error
	select {
	case err := <-finished:
		actualError = err
	case v := <-panicked.Caught():
		actualError = lang.WrapPanic(v)
	}
	if actualError.Error() != "error" {
		t.Errorf("Expected error message 'error', but got '%s'", actualError.Error())
	}

	go func() {
		defer panicked.Recover()
		panic("panicked")
	}()

	select {
	case err := <-finished:
		actualError = err
	case v := <-panicked.Caught():
		actualError = lang.WrapPanic(v)
	}
	panicError, ok := (actualError).(*lang.PanicError)
	if !ok {
		t.Errorf("Expected error is common.PanicError, but got %T", actualError)
	}
	if panicError.Error() != "panicked with panicked" {
		t.Errorf("Expected error message 'panicked with panicked', but got '%s'", panicError.Error())
	}
	if panicError.Origin() != "panicked" {
		t.Errorf("Expected panic 'panicked', but got '%s'", panicError.Origin())
	}
}
