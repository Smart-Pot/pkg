package amqp

import (
	"fmt"

	"github.com/streadway/amqp"
)

type producer struct {
	ch    *amqp.Channel
	queue string
}

func (p *producer) init() error {
	channel := p.ch
	// We create an exahange that will bind to the queue to send and receive messages
	err := channel.ExchangeDeclare(p.queue+exchange, "topic", true, false, false, false, nil)

	if err != nil {
		return err
	}

	return nil
}

func (p *producer) produceStr(s string) error {
	message := amqp.Publishing{
		Body: []byte(s),
	}
	return p.produce(message)

}

func (p *producer) produce(msg amqp.Publishing) error {

	// We publish the message to the exahange we created earlier
	err := p.ch.Publish(p.queue+exchange, "random-key", false, false, msg)
	if err != nil {
		return fmt.Errorf("error publishing a message to the queue: %s", err)
	}
	return nil
}
