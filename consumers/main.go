package consumers

import (
	"github.com/streadway/amqp"
)

func Init() *amqp.Connection {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	return connection
}
