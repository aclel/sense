package main

import (
	"fmt"
	"log"

	mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	opts.SetClientID("go-server")

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	if token := c.Subscribe("test/topic1", 0, testHandler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	for {
	}
}

func testHandler(client *mqtt.Client, msg mqtt.Message) {
	fmt.Println(msg.Topic() + " " + string(msg.Payload()))
}
