package config

import (
	"log"
	"sync"

	"github.com/4klb/coffee-shop-microservices/dish/utils"
)

var (
	once sync.Once
	cfg  *Config
)

type Config struct {
	RabbitMQ struct {
		User     string `env:"RABBITMQ_USER"`
		Password string `env:"RABBITMQ_PASSWORD"`
		Host     string `env:"RABBITMQ_HOST"`
		Port     string `env:"RABBITMQ_PORT"`
		//settings
		ExchangeName string `env:"RABBITMQ_EXCHANGE_NAME"`
		ExchangeType string `env:"RABBITMQ_EXCHANGE_TYPE"`
		RoutingKey   string `env:"RABBITMQ_ROUTING_KEY"`
		QueueName    string `env:"RABBITMQ_QUEUE_NAME"`
	}
	DishServer struct {
		Port    string `env:"DISH_SERVER_PORT"`
		Host    string `env:"DISH_SERVER_HOST"`
		Network string `env:"DISH_SERVER_NETWORK"`
	}
	Postgres struct {
		Port     string `env:"POSTGRES_PORT"`
		User     string `env:"POSTGRES_USER"`
		Password string `env:"POSTGRES_PASSWORD"`
		DB       string `env:"POSTGRES_DB"`
		Host     string `env:"POSTGRES_HOST"`
	}
}

func GetConfig() *Config {
	once.Do(func() {
		config := &Config{}
		if err := utils.LoadConfig(config); err != nil {
			log.Println("Configuration set up failed")
			return
		}
		cfg = config
	})
	return cfg
}
