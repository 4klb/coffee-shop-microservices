package main

import (
	"github.com/4klb/coffee-shop-microservices/gateway/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	server.Run()
}
