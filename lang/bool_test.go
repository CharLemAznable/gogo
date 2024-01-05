package lang_test

import (
	"github.com/CharLemAznable/gogo/lang"
	"testing"
)

func TestToBool(t *testing.T) {
	assertFalse(t, lang.ToBool(""))
	assertTrue(t, lang.ToBool("true"))
	assertTrue(t, lang.ToBool("TRUE"))
	assertTrue(t, lang.ToBool("tRUe"))
	assertTrue(t, lang.ToBool("on"))
	assertTrue(t, lang.ToBool("tRUe"))
	assertTrue(t, lang.ToBool("T"))
	assertFalse(t, lang.ToBool("false"))
	assertFalse(t, lang.ToBool("f"))
	assertFalse(t, lang.ToBool("No"))
	assertFalse(t, lang.ToBool("n"))
	assertTrue(t, lang.ToBool("on"))
	assertTrue(t, lang.ToBool("ON"))
	assertFalse(t, lang.ToBool("off"))
	assertFalse(t, lang.ToBool("oFf"))
	assertTrue(t, lang.ToBool("yes"))
	assertTrue(t, lang.ToBool("Y"))
	assertTrue(t, lang.ToBool("1"))
	assertFalse(t, lang.ToBool("0"))
	assertFalse(t, lang.ToBool("blue"))
	assertFalse(t, lang.ToBool("true "))
	assertFalse(t, lang.ToBool("ono"))
	assertFalse(t, lang.ToBool("oo"))
	assertFalse(t, lang.ToBool("o"))
	assertFalse(t, lang.ToBool("x gti"))
	assertFalse(t, lang.ToBool("x gti "))
}

func assertTrue(t *testing.T, b bool) {
	if !b {
		t.Error("Expected true, but got false")
	}
}

func assertFalse(t *testing.T, b bool) {
	if b {
		t.Error("Expected false, but got true")
	}
}
