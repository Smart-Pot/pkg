/*
Package amqp implements a simple amqp consumer to get message from an queue

Usage:

	package main

	import (
		...
		"fmt"
		"github.com/Smart-Pot/pkg/adapter/amqp"
		...
	)

	func main(){
		...

		consumer,err := amqp.MakeConsumer("queue_name","exchange_name")
		if err != nil {
			// handle error
		}

		msg := consumer.Consume()
		fmt.Println(msg)

		...
	}

*/
package amqp

import (
	"github.com/streadway/amqp"
)

// Consumer represents an AMQP consumer for topic exchange model
type Consumer interface {
	Consume() []byte
}


// MakeConsumer makes consumer with given queue name for the given exchange  
func MakeConsumer(queue, exchange string) (Consumer, error) {
	if !_isSet {
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

	// Make sure that exchange is exist
	err := c.ch.ExchangeDeclare(c.exchange, "topic", true, false, false, false, nil)

	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	_, err = channel.QueueDeclare(c.queue, true, false, false, false, nil)

	if err != nil {
		return err
	}

	// Bind queue and exchange with wildcard '#'
	// For more information : https://www.rabbitmq.com/tutorials/tutorial-five-go.html
	err = channel.QueueBind(c.queue, "#", c.exchange, false, nil)
	if err != nil {
		return err
	}
	
	// Create a channel that is notified when a new message  arrives in the queue
	msgs, err := c.ch.Consume(c.queue, "", true, false, false, false, nil)

	if err != nil {
		return err
	}
	c.msgs = msgs
	return nil
}

// Consume consume a message from consumer's queue and returns it body
func (c *consumer) Consume() []byte {
	return c._consume().Body
}

func (c *consumer) _consume() amqp.Delivery {
	return <-c.msgs
}
