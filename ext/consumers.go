package ext

import (
	"github.com/CharLemAznable/gogo/fn"
	"github.com/CharLemAznable/gogo/lang"
	"sync"
)

type Consumers[T any] interface {
	Accept(T)
	CheckedAccept(T) error
	AppendConsumer(fn.Consumer[T]) Consumers[T]
	RemoveConsumer(fn.Consumer[T]) Consumers[T]
}

type consumers[T any] struct {
	sync.RWMutex
	consumers fn.Consumers[T]
}

func (s *consumers[T]) Accept(t T) {
	s.RLock()
	defer s.RUnlock()
	s.consumers.Accept(t)
}

func (s *consumers[T]) CheckedAccept(t T) error {
	s.RLock()
	defer s.RUnlock()
	return s.consumers.CheckedAccept(t)
}

func (s *consumers[T]) AppendConsumer(consumer fn.Consumer[T]) Consumers[T] {
	s.Lock()
	defer s.Unlock()
	s.consumers = lang.AppendElementUnique[fn.Consumer[T]](s.consumers, consumer)
	return s
}

func (s *consumers[T]) RemoveConsumer(consumer fn.Consumer[T]) Consumers[T] {
	s.Lock()
	defer s.Unlock()
	s.consumers = lang.RemoveElementByValue[fn.Consumer[T]](s.consumers, consumer)
	return s
}

func NewConsumers[T any]() Consumers[T] {
	return &consumers[T]{consumers: make(fn.Consumers[T], 0)}
}
