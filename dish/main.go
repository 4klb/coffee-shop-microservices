package main

import (
	"log"

	"github.com/4klb/coffee-shop-microservices/dish/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	if err := server.InitDishServer(); err != nil {
		log.Fatalf("Couldn't run dish server: %s", err)
	}

	// errC, err := server.Run(rmq)
	// if err != nil {
	// 	log.Fatalf("Couldn't run: %s", err)
	// }

	// if err := <-errC; err != nil {
	// 	log.Fatalf("Error while running: %s", err)
	// }
}
