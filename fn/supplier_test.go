package fn_test

import (
	"errors"
	. "github.com/CharLemAznable/gogo/fn"
	"testing"
)

func TestSupplierOf(t *testing.T) {
	// Test case 1
	fn := func() int {
		return 10
	}
	s := SupplierOf(fn)
	result, err := s.CheckedGet()
	if err != nil {
		t.Error("Expected no error, but got:", err)
	}
	if result != 10 {
		t.Error("Expected 10, but got:", result)
	}

	// Test case 2
	fn = func() int {
		return 0
	}
	s = SupplierOf(fn)
	result, err = s.CheckedGet()
	if err != nil {
		t.Error("Expected no error, but got:", err)
	}
	if result != 0 {
		t.Error("Expected 0, but got:", result)
	}
}

func TestSupplierCast(t *testing.T) {
	// Test case 1
	fn := func() (int, error) {
		return 10, nil
	}
	s := SupplierCast(fn)
	result := s.Get()
	if result != 10 {
		t.Error("Expected 10, but got:", result)
	}

	// Test case 2
	fn = func() (int, error) {
		return 0, nil
	}
	s = SupplierCast(fn)
	result = s.Get()
	if result != 0 {
		t.Error("Expected 0, but got:", result)
	}

	// Test case 3
	fn = func() (int, error) {
		return 0, errors.New("error")
	}
	s = SupplierCast(fn)
	result = s.Get()
	if result != 0 {
		t.Error("Expected 0, but got:", result)
	}
}

func TestConstant(t *testing.T) {
	// Test case 1: Test with integer value
	intSupplier := Zero[int]()
	result := intSupplier.Get()
	if result != 0 {
		t.Errorf("Expected result to be 0, got %v", result)
	}

	// Test case 2: Test with string value
	strSupplier := Constant("Hello")
	result2 := strSupplier.Get()
	if result2 != "Hello" {
		t.Errorf("Expected result to be 'Hello', got %v", result2)
	}

	// Test case 3: Test with custom struct
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "John", Age: 30}
	personSupplier := Constant(person)
	result3 := personSupplier.Get()
	if result3 != person {
		t.Errorf("Expected result to be %+v, got %+v", person, result3)
	}
}
