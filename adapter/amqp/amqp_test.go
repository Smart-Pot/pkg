package amqp

import (
	"fmt"
	"testing"

	"github.com/streadway/amqp"
)

func TestAMQP(t *testing.T) {

	url := "amqp://guest:guest@localhost:5672"

	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(url)
	er(t, err)
	// We close the connection after the operation has completed.
	defer connection.Close()

	ch, err := connection.Channel()
	er(t, err)

	const queue = "tata"

	c := consumer{ch, queue, make(<-chan amqp.Delivery)}
	p := producer{ch, queue}

	if err := p.init(); err != nil {
		t.Error(err)
		t.FailNow()
	}

	if err := c.init(); err != nil {
		t.Error(err)
		t.FailNow()
	}

	const testmsg = "asdasd"

	for i := 0; i < 4; i++ {
		p.produceStr(testmsg)
	}

	k := 0
	for i := 0; i < 4; i++ {
		got := string(c.consume().Body)
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
