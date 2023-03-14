package server

import (
	"log"
	"net"

	"github.com/4klb/coffeetime/dish/config"
	pb "github.com/4klb/coffeetime/proto/dish"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedDishServiceServer
}

func InitServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:"+config.GetConfig().Server.Port)
	if err != nil {
		log.Println(err)
		return
	}

	grpcServer := grpc.NewServer()

	pb.RegisterDishServiceServer(grpcServer, &Server{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
		return
	}
}
