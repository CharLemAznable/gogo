package ext_test

import (
	"github.com/CharLemAznable/gogo/ext"
	"testing"
)

func TestPubSub(t *testing.T) {
	ach := make(chan *MsgA, 1)
	bch := make(chan *MsgB, 1)

	aSub := ext.SubFn(func(msg *MsgA) {
		ach <- msg
	})
	bSub := ext.SubChan(bch)
	subs := ext.JoinSubscribers(aSub, bSub)

	ext.Subscribe("test_topic", subs)

	go func() {
		ext.Publish("test_topic", &MsgA{Content: "MSG_A"})
		ext.Publish("test_topic", &MsgB{Message: "MSG_B"})
	}()

	select {
	case msgA := <-ach:
		if msgA.Content != "MSG_A" {
			t.Errorf("Expected msgA.Content is 'MSG_A', but got '%s'", msgA.Content)
		}
	}
	select {
	case msgB := <-bch:
		if msgB.Message != "MSG_B" {
			t.Errorf("Expected msgB.Message is 'MSG_B', but got '%s'", msgB.Message)
		}
	}

	ext.Unsubscribe("test_topic", subs)
	ext.Unsubscribe("test_topic", bSub)
	ext.Unsubscribe("test_topic", aSub)
}

type MsgA struct {
	Content string
}

type MsgB struct {
	Message string
}
