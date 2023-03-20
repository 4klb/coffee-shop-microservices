package bootstrap

import (
	"github.com/4klb/coffee-shop-microservices/dish/broker"
)

func InitRabbit() {
	broker.GetRabbitConnection()
}
