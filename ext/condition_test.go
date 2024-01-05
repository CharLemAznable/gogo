package ext_test

import (
	"bytes"
	"fmt"
	"github.com/CharLemAznable/gogo/ext"
	"testing"
)

func TestCheckEmptyRun(t *testing.T) {
	emptyCount := 0
	emptyFn := func() {
		emptyCount++
	}
	notEmptyBuf := &bytes.Buffer{}
	notEmptyFn := func(s string) {
		_, _ = fmt.Fprintf(notEmptyBuf, "[%s]", s)
	}

	var str string
	ext.CheckEmptyRun(str, emptyFn, notEmptyFn)
	ext.EmptyThenRun(str, emptyFn)
	ext.NotEmptyThenRun(str, notEmptyFn)
	str = "abc"
	ext.CheckEmptyRun(str, emptyFn, notEmptyFn)
	ext.EmptyThenRun(str, emptyFn)
	ext.NotEmptyThenRun(str, notEmptyFn)

	if 2 != emptyCount {
		t.Errorf("Expected emptyCount is 2, but got: %d", emptyCount)
	}
	if "[abc][abc]" != notEmptyBuf.String() {
		t.Errorf("Expected notEmptyBuf is '[abc][abc]', but got: %s", notEmptyBuf.String())
	}
}

func TestCheckEmpty(t *testing.T) {
	emptyFn := func() string {
		return "def"
	}
	notEmptyFn := func(s string) string {
		return fmt.Sprintf("[%s]", s)
	}

	var str string
	ret := ext.CheckEmpty(str, emptyFn, notEmptyFn)
	if "def" != ret {
		t.Errorf("Expected ret is 'def', but got: %s", ret)
	}
	ret = ext.EmptyThen(str, emptyFn)
	if "def" != ret {
		t.Errorf("Expected ret is 'def', but got: %s", ret)
	}
	ret = ext.NotEmptyThen(str, notEmptyFn)
	if "" != ret {
		t.Errorf("Expected ret is '', but got: %s", ret)
	}
	str = "abc"
	ret = ext.CheckEmpty(str, emptyFn, notEmptyFn)
	if "[abc]" != ret {
		t.Errorf("Expected ret is '[abc]', but got: %s", ret)
	}
	ret = ext.EmptyThen(str, emptyFn)
	if "abc" != ret {
		t.Errorf("Expected ret is 'abc', but got: %s", ret)
	}
	ret = ext.NotEmptyThen(str, notEmptyFn)
	if "[abc]" != ret {
		t.Errorf("Expected ret is '[abc]', but got: %s", ret)
	}
}
