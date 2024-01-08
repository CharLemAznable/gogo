package fn_test

import (
	"bytes"
	"errors"
	. "github.com/CharLemAznable/gogo/fn"
	"strconv"
	"testing"
)

func TestFunctionOf(t *testing.T) {
	// Test case 1: Test with a function that takes a string and returns an integer
	fn := FunctionOf(func(t string) int {
		return len(t)
	})

	result, err := fn.CheckedApply("Hello")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 2: Test with a function that takes an integer and returns a string
	fn2 := FunctionOf(func(t int) string {
		if t%2 == 0 {
			return "even"
		}
		return "odd"
	})

	resultStr, err := fn2.CheckedApply(7)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	expectedStr := "odd"
	if resultStr != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, resultStr)
	}
}

func TestFunctionCast(t *testing.T) {
	// Test case 1: Test with a function that returns a string
	fn := FunctionCast(func(t string) (string, error) {
		return "Hello, " + t, nil
	})

	result := fn.Apply("World")
	expected := "Hello, World"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Test case 2: Test with a function that returns an integer
	fn2 := FunctionCast(func(t int) (int, error) {
		return t * 2, nil
	})

	result2 := fn2.Apply(5)
	expectedInt := 10
	if result2 != expectedInt {
		t.Errorf("Expected %d, but got %d", expectedInt, result2)
	}

	// Test case 3: Test with a function that returns with error
	fn3 := FunctionCast(func(s string) (int, error) {
		return 0, errors.New("error")
	})

	result3 := fn3.Apply("5")
	expectedInt = 0
	if result3 != expectedInt {
		t.Errorf("Expected %d, but got %d", expectedInt, result2)
	}
}

func TestComposeFunction(t *testing.T) {
	fn1 := func(s string) (int, error) {
		return strconv.Atoi(s)
	}
	fn2 := func(i int) (string, error) {
		buf := &bytes.Buffer{}
		for index := 0; index < i; index++ {
			buf.WriteString("*")
		}
		return buf.String(), nil
	}

	composeFunction := ComposeFunction(FunctionCast(fn1), FunctionCast(fn2))
	ret := composeFunction.Apply("a")
	if ret != "" {
		t.Errorf("Expected '', but got %s", ret)
	}
	ret, err := composeFunction.CheckedApply("a")
	if ret != "" {
		t.Errorf("Expected '', but got %s", ret)
	}
	if err == nil {
		t.Error("Expected get error, but not")
	}
	ret, err = composeFunction.CheckedApply("2")
	if ret != "**" {
		t.Errorf("Expected '**', but got %s", ret)
	}
	if err != nil {
		t.Errorf("Expected no error, but got '%s'", err.Error())
	}
}

func TestIdentity(t *testing.T) {
	// Test case 1
	result := Identity[int]().Apply(5)
	if result != 5 {
		t.Errorf("Expected result to be 5, but got: %v", result)
	}

	// Test case 2
	result2, err := Identity[string]().CheckedApply("hello")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if result2 != "hello" {
		t.Errorf("Expected result to be 'hello', but got: %v", result2)
	}
}

func TestYCombinator(t *testing.T) {
	// Test case 1
	fab := func(g func(int) int) func(int) int {
		return func(n int) int {
			if n == 0 {
				return 1
			}
			return n * g(n-1)
		}
	}
	result := YCombinator(fab).Apply(5)
	if result != 120 {
		t.Errorf("Expected result to be 120, but got: %v", result)
	}

	// Test case 2
	fib := func(g func(int) int) func(int) int {
		return func(i int) int {
			if i <= 2 {
				return 1
			}
			return g(i-1) + g(i-2)
		}
	}
	result2 := YCombinator(fib).Apply(10)
	if result2 != 55 {
		t.Errorf("Expected result to be 55, but got: %v", result2)
	}
}
