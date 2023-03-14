package bootstrap

import (
	"github.com/4klb/coffeetime/dish/broker"
)

func InitRabbit() {
	broker.GetRabbitConnection()
}
