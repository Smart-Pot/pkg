package amqp

import (
	"fmt"
	"testing"
)

func TestAMQP(t *testing.T) {

	url := "amqp://guest:guest@localhost:5672"

	if err := Set(url); err != nil {
		t.Error(err)
		t.FailNow()
	}

	const queue = "tata"
	const exchange = "testos"

	c, err := MakeConsumer(queue, exchange)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	p, err := MakeProducer(exchange)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	const testmsg = "asdasd"

	for i := 0; i < 4; i++ {
		p.Produce([]byte(testmsg))
	}

	k := 0
	for i := 0; i < 4; i++ {
		got := string(c.Consume())
		if testmsg == got {
			k++
		} else {
			t.Error(fmt.Sprintf("want %s but got %s", testmsg, got))
			t.FailNow()
		}
	}

	if k != 4 {
		t.Error(fmt.Sprintf("K is not 4 k: %d", k))
		t.FailNow()
	}

}

func er(t *testing.T, err error) {

}
