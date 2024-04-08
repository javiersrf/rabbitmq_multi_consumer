package main

import (
	"fmt"

	"github.com/jsrf-consulting/easy_diet_go_consumer/consumers"
	"github.com/jsrf-consulting/easy_diet_go_consumer/handlers"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ in Golang: Getting started tutorial")
	conn := consumers.Init()
	defer conn.Close()

	channel, err := conn.Channel()
	handlers.HandleError(err, "Error on opening Channel")

	defer channel.Close()
	fmt.Println("Successfully connected to RabbitMQ instance")

	forever := make(chan bool)
	worker1 := consumers.Consumer{
		Name:    "test",
		AutoAck: true,
		Callback: func(msg *amqp.Delivery) {
			fmt.Printf("Received Message : %s\n", msg.Body)
		},
	}
	worker2 := consumers.Consumer{
		Name:    "worker2",
		AutoAck: true,
		Callback: func(msg *amqp.Delivery) {
			fmt.Printf("Received Message : %s\n", msg.Body)
		},
	}

	go worker1.StartConsumer(channel)
	go worker2.StartConsumer(channel)
	<-forever

}
