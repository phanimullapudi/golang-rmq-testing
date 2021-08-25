package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Service interface {
	Connect() error
}

type RabbitMQ struct {
	Conn *amqp.Connection
}

func (rmq *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RMQ....")
	var err error
	rmq.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	return nil
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
