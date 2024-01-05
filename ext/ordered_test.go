package ext_test

import (
	"bytes"
	"fmt"
	"github.com/CharLemAznable/gogo/ext"
	"testing"
)

func TestOrdered(t *testing.T) {
	orderedSlice := ext.JoinOrdered(
		&OrderItem{Name: "AAA", order: "3"},
		&OrderItem{Name: "BBB", order: "1"},
		&OrderItem{Name: "CCC", order: "2"})
	orderedSlice.Sort()

	buf := &bytes.Buffer{}
	for _, item := range orderedSlice {
		_, _ = fmt.Fprintf(buf, "[%s]", item)
	}
	if "[BBB][CCC][AAA]" != buf.String() {
		t.Errorf("Expected is '[BBB][CCC][AAA]', but got '%s'", buf.String())
	}
}

type OrderItem struct {
	Name  string
	order string
}

func (i *OrderItem) Order() string {
	return i.order
}

func (i *OrderItem) String() string {
	return i.Name
}
