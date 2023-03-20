package main

import (
	"github.com/4klb/coffee-shop-microservices/dish/bootstrap"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bootstrap.InitRabbit()
}
