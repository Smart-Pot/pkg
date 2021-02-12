/*
Package amqp implements a simple library that creates and abstraction layer for consumer producer communication

In general, an AMQP server has several exchange models.
But we use only the topic exchange model to either satisfy the needs of the Smart-Pot project
and make it simple and testable as possible as we can.

Usage:

	package main

	import (
		...
		"github.com/Smart-Pot/pkg/adapter/amqp"
		...
	)

	func main(){
		...
		url := ...
		if err := amqp.Set(url); err != nil {
			// Handle error
		}
	}


*/
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
