package amqp

import (
	"github.com/streadway/amqp"
)

type Consumer interface {
	Consume() []byte
}

func MakeConsumer(queue, exchange string) (Consumer, error) {
	if !isSet {
		return nil, ErrNotSet
	}
	c := &consumer{
		ch:       _channel,
		queue:    queue,
		exchange: exchange,
		msgs:     make(<-chan amqp.Delivery),
	}

	if err := c.init(); err != nil {
		return nil, err
	}
	return c, nil
}

type consumer struct {
	ch       *amqp.Channel
	queue    string
	exchange string
	msgs     <-chan amqp.Delivery
}

func (c *consumer) init() error {
	channel := c.ch

	err := c.ch.ExchangeDeclare(c.exchange, "topic", true, false, false, false, nil)

	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	// We create a queue named Test
	_, err = channel.QueueDeclare(c.queue, true, false, false, false, nil)

	if err != nil {
		return err
	}

	// We bind the queue to the exchange to send and receive data from the queue
	err = channel.QueueBind(c.queue, "#", c.exchange, false, nil)
	if err != nil {
		return err
	}
	// We consume data from the queue named Test using the channel we created in go.
	msgs, err := c.ch.Consume(c.queue, "", true, false, false, false, nil)

	if err != nil {
		return err
	}
	c.msgs = msgs
	return nil
}

func (c *consumer) Consume() []byte {
	return c._consume().Body
}

func (c *consumer) _consume() amqp.Delivery {
	return <-c.msgs
}
