package ext

import (
	"github.com/CharLemAznable/gogo/fn"
	"github.com/CharLemAznable/gogo/lang"
	"sync"
)

type PubSub interface {
	Subscribe(string, Subscriber)
	Unsubscribe(string, Subscriber)
	Publish(string, any)
}

func NewPubSub() PubSub {
	return &hub{subscribers: make(map[string][]Subscriber)}
}

type hub struct {
	sync.RWMutex
	subscribers map[string][]Subscriber
}

func (h *hub) Subscribe(topic string, subscriber Subscriber) {
	h.Lock()
	defer h.Unlock()
	h.subscribers[topic] = lang.AppendElementUnique(h.subscribers[topic], subscriber)
}

func (h *hub) Unsubscribe(topic string, subscriber Subscriber) {
	h.Lock()
	defer h.Unlock()
	if subscribers, ok := h.subscribers[topic]; ok {
		h.subscribers[topic] = lang.RemoveElementByValue(subscribers, subscriber)
	}
}

func (h *hub) Publish(topic string, msg any) {
	h.RLock()
	defer h.RUnlock()
	if subscribers, ok := h.subscribers[topic]; ok {
		for _, sub := range subscribers {
			go sub.Subscribe(msg)
		}
	}
}

type Subscriber interface {
	Subscribe(any)
}

type SubscribeConsumer[T any] struct {
	fn.Consumer[T]
}

func (s SubscribeConsumer[T]) Subscribe(msg any) {
	if message, ok := msg.(T); ok {
		s.Consumer.Accept(message)
	}
}

func SubConsumer[T any](c fn.Consumer[T]) Subscriber {
	return SubscribeConsumer[T]{c}
}

func SubFn[T any](f func(T)) Subscriber {
	return SubConsumer(fn.ConsumerOf(f))
}

func SubChan[T any](ch chan T) Subscriber {
	return SubConsumer(fn.ConsumerChan(ch))
}

type Subscribers []Subscriber

func (s Subscribers) Subscribe(msg any) {
	for _, sub := range s {
		go sub.Subscribe(msg)
	}
}

func JoinSubscribers(subscribers ...Subscriber) Subscriber {
	return Subscribers(subscribers)
}

var globalPubSub = NewPubSub()

func Subscribe(topic string, subscriber Subscriber) {
	globalPubSub.Subscribe(topic, subscriber)
}

func Unsubscribe(topic string, subscriber Subscriber) {
	globalPubSub.Unsubscribe(topic, subscriber)
}

func Publish(topic string, msg any) {
	globalPubSub.Publish(topic, msg)
}
