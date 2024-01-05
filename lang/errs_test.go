package lang_test

import (
	"errors"
	"github.com/CharLemAznable/gogo/lang"
	"testing"
)

func TestDefaultErrorMsg(t *testing.T) {
	// Test case 1: err is nil, defMsg is not empty
	err1 := lang.DefaultErrorMsg(nil, "default message")
	if err1 == nil {
		t.Error("Expected an error, but got nil")
	} else if err1.Error() != "default message" {
		t.Errorf("Expected error message 'default message', but got '%s'", err1.Error())
	}

	// Test case 2: err is not nil, defMsg is empty
	err2 := errors.New("original error")
	err2 = lang.DefaultErrorMsg(err2, "")
	if err2 != err2 {
		t.Error("Expected the same error, but got a different error")
	}

	// Test case 3: err is not nil, defMsg is not empty
	err3 := errors.New("original error")
	err3 = lang.DefaultErrorMsg(err3, "default message")
	if err3 == nil {
		t.Error("Expected an error, but got nil")
	} else if err3.Error() != "original error" {
		t.Errorf("Expected error message 'original error', but got '%s'", err3.Error())
	}
}

func TestMultiError(t *testing.T) {
	me := lang.MultiError{}
	me.Append(nil)
	if me.Error() != "" {
		t.Errorf("Expected empty error, but got '%s'", me.Error())
	}
	if me.MaybeUnwrap() != nil {
		t.Errorf("Expected empty error, but got '%v'", me.MaybeUnwrap())
	}

	me.Append(errors.New("error"))
	if me.Error() != "1 error(s) occurred:\n* error" {
		t.Errorf("Expected 1 error, but got '%s'", me.Error())
	}
	if me.MaybeUnwrap().Error() != "error" {
		t.Errorf("Expected 'error', but got '%v'", me.MaybeUnwrap())
	}

	me.Append(errors.New("error2"))
	if me.Error() != "2 error(s) occurred:\n* error\n* error2" {
		t.Errorf("Expected 2 errors, but got '%s'", me.Error())
	}
	if me.MaybeUnwrap().Error() != me.Error() {
		t.Errorf("Expected Unwrap self, but got '%v'", me.MaybeUnwrap())
	}
}
