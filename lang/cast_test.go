package lang_test

import (
	"fmt"
	"github.com/CharLemAznable/gogo/lang"
	"testing"
)

func TestCast(t *testing.T) {
	t.Run("Test with valid type", func(t *testing.T) {
		val := "test"
		result, err := lang.CastOrZero[string](val)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if result != val {
			t.Errorf("Expected result to be %s, but got %s", val, result)
		}
		result = lang.CastQuietly[string](val)
		if result != val {
			t.Errorf("Expected result to be %s, but got %s", val, result)
		}
	})

	t.Run("Test with invalid type", func(t *testing.T) {
		val := 123
		_, err := lang.CastOrZero[string](val)
		if err == nil {
			t.Error("Expected an error, but got none")
		}
		expectedErr := fmt.Sprintf("%#v type mismatch %T", val, "")
		if err.Error() != expectedErr {
			t.Errorf("Expected error message to be %s, but got %s", expectedErr, err.Error())
		}
		result := lang.CastQuietly[string](val)
		if result != "" {
			t.Errorf("Expected result to be %s, but got %s", "", result)
		}
	})

	t.Run("Test with nil value", func(t *testing.T) {
		var val interface{} = nil
		result, err := lang.CastOrZero[string](val)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if result != "" {
			t.Errorf("Expected result to be %s, but got %s", "", result)
		}
		result = lang.Zero[string]()
		if result != "" {
			t.Errorf("Expected result to be %s, but got %s", "", result)
		}
	})
}
