package config

import (
	"log"
	"sync"

	"github.com/4klb/coffeetime/dish/utils"
)

var (
	once sync.Once
	cfg  *Config
)

type Config struct {
	Server struct {
		Port string `env:"SERVER_PORT"`
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
