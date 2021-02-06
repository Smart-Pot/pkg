package amqp

import (
	"fmt"

	"github.com/streadway/amqp"
)

type consumer struct {
	ch    *amqp.Channel
	queue string
	msgs  <-chan amqp.Delivery
}

func (c *consumer) init() error {
	channel := c.ch
	var err error
	if err != nil {
		panic("ABOV" + err.Error())
	}
	// We create a queue named Test
	_, err = channel.QueueDeclare(c.queue, true, false, false, false, nil)

	if err != nil {
		fmt.Println("1")
		return err
	}

	// We bind the queue to the exchange to send and receive data from the queue
	err = channel.QueueBind(c.queue, "#", c.queue+exchange, false, nil)
	if err != nil {
		fmt.Println("2")
		return err
	}
	// We consume data from the queue named Test using the channel we created in go.
	msgs, err := c.ch.Consume(c.queue, "", false, false, false, false, nil)

	if err != nil {
		fmt.Println("3")

		return err
	}
	c.msgs = msgs
	return nil
}

func (c *consumer) consume() amqp.Delivery {
	return <-c.msgs
}
