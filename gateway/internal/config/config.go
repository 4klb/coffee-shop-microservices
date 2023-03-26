package config

import (
	"log"
	"sync"

	"github.com/4klb/coffee-shop-microservices/utils"
)

var (
	once sync.Once
	cfg  *Config
)

//Config ..
type Config struct {
	Api struct {
		Port string `env:"PORT_API"`
	}
	MiddlewareAuth struct {
		Login    string `env:"LOGIN_AUTH"`
		Password string `env:"PASSWORD_AUTH"`
	}
}

//GetConfig ..
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
