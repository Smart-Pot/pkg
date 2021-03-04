package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

/* HANDLERS */

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    fmt.Println("Connected")
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    fmt.Printf("Connect lost: %v", err)
}



func Connect(address,username,password,clientID string) (Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(address)
    opts.SetClientID(clientID)
	opts.SetUsername(username)
    opts.SetPassword(password)
    opts.SetDefaultPublishHandler(MessagePubHandler)
    opts.OnConnect = ConnectHandler
    opts.OnConnectionLost = ConnectLostHandler
    cl := mqtt.NewClient(opts)
	if t := cl.Connect(); t.Wait() && t.Error() != nil {
		return nil,t.Error()
	}
	return &client{cl},nil
}





