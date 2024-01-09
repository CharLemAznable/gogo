package fn_test

import (
	"errors"
	. "github.com/CharLemAznable/gogo/fn"
	"github.com/CharLemAznable/gogo/lang"
	"testing"
)

func TestMustRun(t *testing.T) {
	err := errors.New("must error")
	fn1 := MustRun(RunnableCast(func() error {
		return err
	}))
	fn2 := MustRun(RunnableOf(func() {
		panic(err)
	}))

	err1 := fn1.CheckedRun()
	if !lang.Equal(err, err1) {
		t.Errorf("Expected err1 is err, but got '%v'", err1)
	}
	err2 := fn2.CheckedRun()
	if !lang.Equal(err, err2.(*lang.PanicError).Origin()) {
		t.Errorf("Expected err2.Origin is err, but got '%v'", err2)
	}

	fn2.Run()
}

func TestMustGet(t *testing.T) {
	err := errors.New("must error")
	fn1 := MustGet(SupplierCast(func() (int, error) {
		return 0, err
	}))
	fn2 := MustGet(SupplierCast(func() (int, error) {
		panic(err)
	}))

	_, err1 := fn1.CheckedGet()
	if !lang.Equal(err, err1) {
		t.Errorf("Expected err1 is err, but got '%v'", err1)
	}
	_, err2 := fn2.CheckedGet()
	if !lang.Equal(err, err2.(*lang.PanicError).Origin()) {
		t.Errorf("Expected err2.Origin is err, but got '%v'", err2)
	}

	v := fn2.Get()
	if v != 0 {
		t.Errorf("Expected get 0, but got '%d'", v)
	}
}

func TestMustAccept(t *testing.T) {
	err := errors.New("must error")
	fn1 := MustAccept(ConsumerCast(func(_ string) error {
		return err
	}))
	fn2 := MustAccept(ConsumerCast(func(_ string) error {
		panic(err)
	}))

	err1 := fn1.CheckedAccept("")
	if !lang.Equal(err, err1) {
		t.Errorf("Expected err1 is err, but got '%v'", err1)
	}
	err2 := fn2.CheckedAccept("")
	if !lang.Equal(err, err2.(*lang.PanicError).Origin()) {
		t.Errorf("Expected err2.Origin is err, but got '%v'", err2)
	}

	fn2.Accept("")
}

func TestMustApply(t *testing.T) {
	err := errors.New("must error")
	fn1 := MustApply(FunctionCast(func(_ string) (int, error) {
		return 0, err
	}))
	fn2 := MustApply(FunctionCast(func(_ string) (int, error) {
		panic(err)
	}))

	_, err1 := fn1.CheckedApply("")
	if !lang.Equal(err, err1) {
		t.Errorf("Expected err1 is err, but got '%v'", err1)
	}
	_, err2 := fn2.CheckedApply("")
	if !lang.Equal(err, err2.(*lang.PanicError).Origin()) {
		t.Errorf("Expected err2.Origin is err, but got '%v'", err2)
	}

	v := fn2.Apply("")
	if v != 0 {
		t.Errorf("Expected get 0, but got '%d'", v)
	}
}
