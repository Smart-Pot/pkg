package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Client interface {
	Subscribe(topic string)  (Consumer,error)
	Publish(topic,message string) error
}

type client struct {
	cli mqtt.Client
}


func (c *client) Subscribe(topic string) (Consumer,error) {
 	
	ch := make(chan Message)
	t := c.cli.Subscribe(topic,1,func(c mqtt.Client, m mqtt.Message) {
		ch <- m
	})
	<-t.Done()
	if err := t.Error(); err != nil {
		return nil ,err
	}
	return &consumer{ch},nil
}


func (c *client) Publish(topic,message string) error {
	t := c.cli.Publish(topic,0,false,message)
	<-t.Done()
	return t.Error()
}


type Message = mqtt.Message

type Consumer interface {
	Consume() Message
}

type consumer struct {
	ch chan Message
}


func (c *consumer) Consume() Message {
	return <-c.ch
}