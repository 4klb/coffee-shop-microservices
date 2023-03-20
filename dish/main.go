package main

import (
	"github.com/4klb/coffeetime/dish/bootstrap"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bootstrap.InitRabbit()
}
