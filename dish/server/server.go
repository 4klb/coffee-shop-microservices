package server

import (
	"fmt"
	"log"
	"net"

	"github.com/4klb/coffee-shop-microservices/dish/broker"
	"github.com/4klb/coffee-shop-microservices/dish/config"
	pb "github.com/4klb/coffee-shop-microservices/proto/dish"
	"google.golang.org/grpc"
)

type ServerDish struct {
	pb.UnimplementedDishServiceServer
}

func Run() {
	var consumer broker.Consumer

	go func() {
		if err := InitDishServer(); err != nil {
			log.Fatalf("Couldn't run dish server: %s", err)
		}
	}()

	rmq, err := broker.GetRabbitConnection()
	if err != nil {
		log.Fatalf("broker.GetRabbitConnection: %s", err)
	}

	errC := make(<-chan error, 1)

	go func() {
		errC, err = consumer.RunConsumer(rmq)
		if err != nil {
			log.Fatalf("Couldn't run: %s", err)
		}

	}()
	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func InitDishServer() error {
	log.Printf("Dish server is running...")

	listener, err := net.Listen(
		config.GetConfig().DishServer.Network,
		config.GetConfig().DishServer.Host+config.GetConfig().DishServer.Port)
	if err != nil {
		return fmt.Errorf("net.Listen %w", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterDishServiceServer(grpcServer, &ServerDish{})

	if err = grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("grpcServer.Serve %w", err)
	}

	return nil
}
