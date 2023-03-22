package server

import (
	"fmt"
	"net"

	"github.com/4klb/coffee-shop-microservices/dish/config"
	pb "github.com/4klb/coffee-shop-microservices/proto/dish"
	"google.golang.org/grpc"
)

type ServerDish struct {
	pb.UnimplementedDishServiceServer
}

func InitDishServer() error {
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
