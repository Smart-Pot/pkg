/*
Package amqp implements a simple amqp producer to share message to any exchange

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

		producer,err := amqp.MakeProducer("exchange_name")
		if err != nil {
			// handle error
		}

		if err = producer.Produce(); err != nil {
			// handle err
		}
		...
	}

*/
package amqp

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Producer represent an AMQP producer for topic exchange model
type Producer interface {
	Produce([]byte) error
}

// MakeProducer creates a producer for the given exchange
func MakeProducer(exchange string) (Producer, error) {
	if !_isSet {
		return nil, ErrNotSet
	}
	p := &producer{
		ch:       _channel,
		exchange: exchange,
	}

	if err := p.init(); err != nil {
		return nil, err
	}
	return p, nil
}

type producer struct {
	ch       *amqp.Channel
	exchange string
}

func (p *producer) init() error {
	// Make sure that exchange is exist
	return p.ch.ExchangeDeclare(p.exchange, "topic", true, false, false, false, nil)
}
// Produce publish an message to the exchange of the producer
func (p *producer) Produce(b []byte) error {
	message := amqp.Publishing{
		Body: b,
	}
	return p._produce(message)

}

func (p *producer) _produce(msg amqp.Publishing) error {

	// Publish message to exchange of the producer
	err := p.ch.Publish(p.exchange, "random-key", false, false, msg)
	if err != nil {
		return fmt.Errorf("error publishing a message to the queue: %s", err)
	}
	return nil
}
