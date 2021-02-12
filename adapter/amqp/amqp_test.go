package amqp_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Smart-Pot/pkg/adapter/amqp"
	"github.com/stretchr/testify/assert"
)



const (
	_msgCount = 4
	_queueName = "test_queue"
	_exchangeName = "test_exchange"
	// The url that an AMQP server running on it
	_url = "amqp://guest:guest@localhost:5672"
	
)
var (
	_testMessage = []byte("test_message")
 	_done = make(chan bool,1)
)


// AMQP testing needs a running AMQP server on given url
// Before starting the test, make sure the server is ready to recieve and send messages 
 func TestMain(m *testing.M) {
	// Set RabbitMQ connection
	if err := amqp.Set(_url); err != nil {
		panic(fmt.Errorf("test main: %s",err))
	}
	
	c := m.Run()

	// Wait for consumer test
	<-_done

	os.Exit(c)
}

func TestConsumer(t *testing.T) {
	c, err := amqp.MakeConsumer(_queueName, _exchangeName)
	assert.Nil(t,err)

	go func(){
		for i := 0; i < _msgCount; i++ {
			msg := c.Consume()
			assert.Equal(t,_testMessage,msg)
		}
		_done<- true
	}()
}


func TestProducer(t *testing.T) {
	p,err := amqp.MakeProducer(_exchangeName)
	assert.Nil(t,err)
	for i := 0; i < _msgCount; i++ {
		err = p.Produce(_testMessage)
		assert.Nil(t,err)
	} 
}

