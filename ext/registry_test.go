package ext_test

import (
	"github.com/CharLemAznable/gogo/ext"
	"testing"
)

type RegItem struct {
	Name string
}

func TestSimpleRegistry(t *testing.T) {
	registry := ext.NewSimpleRegistry[*RegItem]()

	err := registry.Register("a", nil)
	if "RegistryError: register item is zero value" != err.Error() {
		t.Errorf("Unexpected err: %s", err.Error())
	}
	err = registry.Register("a", &RegItem{Name: "A"})
	if err != nil {
		t.Errorf("Unexpected err: %s", err.Error())
	}
	err = registry.Register("a", &RegItem{Name: "B"})
	if "RegistryError: multiple registrations for \"a\"" != err.Error() {
		t.Errorf("Unexpected err: %s", err.Error())
	}

	item, err := registry.Get("a")
	if "A" != item.Name {
		t.Errorf("Unexpected get: %s", item.Name)
	}
	if err != nil {
		t.Errorf("Unexpected err: %s", err.Error())
	}
	item, err = registry.Get("b")
	if item != nil {
		t.Errorf("Unexpected get: %s", item.Name)
	}
	if "RegistryError: none registrations for \"b\"" != err.Error() {
		t.Errorf("Unexpected err: %s", err.Error())
	}
}

func TestDefaultRegistry(t *testing.T) {
	registry := ext.NewDefaultRegistry("", &RegItem{Name: "DEF"})

	err := registry.Register("a", nil)
	if "RegistryError: register item is zero value" != err.Error() {
		t.Errorf("Unexpected err: %s", err.Error())
	}
	err = registry.Register("", &RegItem{Name: "A"})
	if "RegistryError: register item name is illegal" != err.Error() {
		t.Errorf("Unexpected err: %s", err.Error())
	}
	err = registry.Register("a", &RegItem{Name: "A"})
	if err != nil {
		t.Errorf("Unexpected err: %s", err.Error())
	}
	err = registry.Register("a", &RegItem{Name: "B"})
	if "RegistryError: multiple registrations for \"a\"" != err.Error() {
		t.Errorf("Unexpected err: %s", err.Error())
	}

	item, err := registry.Get("a")
	if "A" != item.Name {
		t.Errorf("Unexpected get: %s", item.Name)
	}
	if err != nil {
		t.Errorf("Unexpected err: %s", err.Error())
	}
	item, err = registry.Get("")
	if "DEF" != item.Name {
		t.Errorf("Unexpected get: %s", item.Name)
	}
	if err != nil {
		t.Errorf("Unexpected err: %s", err.Error())
	}
	item, err = registry.Get("b")
	if "DEF" != item.Name {
		t.Errorf("Unexpected get: %s", item.Name)
	}
	if err != nil {
		t.Errorf("Unexpected err: %s", err.Error())
	}
}
