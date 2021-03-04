package mqtt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



const (
	_address = "mqtt://localhost:1883"
)

func TestConnect(t *testing.T) {
	c,err := Connect(_address,"","","test-client")
	assert.Nil(t,err)
	assert.NotNil(t,c)

	const (
		topic = "test/example"
		message = "hello message"
	)
	done := make(chan interface{})
	t.Run("Subscribe",func(t *testing.T) {
		cs, err := c.Subscribe(topic)
		assert.Nil(t,err)
		go func() {
			m := cs.Consume()
			assert.Equal(t,message,string(m.Payload()))
			done<-true
		}()
	})
	t.Run("Publish", func(t *testing.T) {
		assert.Nil(t,c.Publish(topic,message))
	})
	<- done
}




