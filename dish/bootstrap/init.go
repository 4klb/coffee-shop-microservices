package bootstrap

import (
	"log"

	"github.com/4klb/coffee-shop-microservices/dish/broker"
)

func InitRabbit() *broker.RabbitMQ {
	rmq, err := broker.GetRabbitConnection()
	if err != nil {
		log.Panic(err)
	}
	return rmq
}
