package main

import (
	"github.com/4klb/coffee-shop-microservices/dish/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	server.Run()
}
