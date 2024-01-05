package ext

import (
	"fmt"
	"github.com/CharLemAznable/gogo/lang"
	"sync"
)

type Registry[T any] interface {
	Register(string, T) error
	Get(string) (T, error)
}

func NewSimpleRegistry[T any]() Registry[T] {
	return &simpleRegistry[T]{table: make(map[string]T)}
}

type simpleRegistry[T any] struct {
	sync.RWMutex
	table map[string]T
}

func (r *simpleRegistry[T]) Register(name string, item T) error {
	r.Lock()
	defer r.Unlock()
	if lang.IsZero(item) {
		return &RegistryError{message: "register item is zero value"}
	}
	if _, exist := r.table[name]; exist {
		return &RegistryError{message: "multiple registrations for \"" + name + "\""}
	}
	r.table[name] = item
	return nil
}

func (r *simpleRegistry[T]) Get(name string) (T, error) {
	r.RLock()
	defer r.RUnlock()
	if item, exist := r.table[name]; exist {
		return item, nil
	}
	return lang.Zero[T](), &RegistryError{message: "none registrations for \"" + name + "\""}
}

func NewDefaultRegistry[T any](defName string, defItem T) Registry[T] {
	return &defaultRegistry[T]{table: make(map[string]T), defName: defName, defItem: defItem}
}

type defaultRegistry[T any] struct {
	sync.RWMutex
	table   map[string]T
	defName string
	defItem T
}

func (r *defaultRegistry[T]) Register(name string, item T) error {
	r.Lock()
	defer r.Unlock()
	if lang.IsZero(item) {
		return &RegistryError{message: "register item is zero value"}
	}
	if name == r.defName {
		return &RegistryError{message: "register item name is illegal"}
	}
	if _, exist := r.table[name]; exist {
		return &RegistryError{message: "multiple registrations for \"" + name + "\""}
	}
	r.table[name] = item
	return nil
}

func (r *defaultRegistry[T]) Get(name string) (T, error) {
	r.RLock()
	defer r.RUnlock()
	if name == r.defName {
		return r.defItem, nil
	}
	if item, exist := r.table[name]; exist {
		return item, nil
	}
	return r.defItem, nil
}

type RegistryError struct {
	message string
}

func (e *RegistryError) Error() string {
	return fmt.Sprintf("RegistryError: %s", e.message)
}
