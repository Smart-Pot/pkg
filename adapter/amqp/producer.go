package amqp

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Producer interface {
	Produce([]byte) error
}

func MakeProducer(exchange string) (Producer, error) {
	if !isSet {
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
	// We create an exahange that will bind to the queue to send and receive messages
	return p.ch.ExchangeDeclare(p.exchange, "topic", true, false, false, false, nil)
}

func (p *producer) Produce(b []byte) error {
	message := amqp.Publishing{
		Body: b,
	}
	return p._produce(message)

}

func (p *producer) _produce(msg amqp.Publishing) error {

	// We publish the message to the exahange we created earlier
	err := p.ch.Publish(p.exchange, "random-key", false, false, msg)
	if err != nil {
		return fmt.Errorf("error publishing a message to the queue: %s", err)
	}
	return nil
}
