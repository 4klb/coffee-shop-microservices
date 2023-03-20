package broker

import (
	"fmt"
	"log"

	"github.com/4klb/coffeetime/dish/config"
	"github.com/streadway/amqp"
)

func GetRabbitConnection() (*amqp.Connection, error) {
	uri := fmt.Sprintf("amqp://%v:%v@%v:%v/", config.GetConfig().RabbitMQ.User, config.GetConfig().RabbitMQ.Password, config.GetConfig().RabbitMQ.Host, config.GetConfig().RabbitMQ.Port)

	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	fmt.Println("Successfully connected to RabbitMQ")

	return conn, nil
}

func GetExchange() {}

func PutDishIntoQueue() {}
