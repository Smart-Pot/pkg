package amqp

import (
	"errors"

	"github.com/streadway/amqp"
)

var _channel *amqp.Channel

var isSet bool = false

var (
	ErrNotSet = errors.New("amqp adapter is not set")
)

func Set(url string) error {

	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(url)

	if err != nil {
		return err
	}
	_channel, err = connection.Channel()
	if err != nil {
		return err
	}
	isSet = true
	return nil
}

func Close() error {
	return _channel.Close()
}
