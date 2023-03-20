package server

import (
	"context"

	"github.com/4klb/coffee-shop-microservices/dish/handler"
	"github.com/4klb/coffee-shop-microservices/proto/dish"
	pb "github.com/4klb/coffee-shop-microservices/proto/dish"
)

// AddDish implements dish.DishServiceServer
func (s *Server) AddDish(ctx context.Context, in *pb.AddDishReq) (*pb.Empty, error) {
	return &dish.Empty{}, handler.Add(ctx, in)
}

// Dishes implements dish.DishServiceServer
func (s *Server) AllDishes(ctx context.Context, in *pb.Empty) (*pb.DishesResp, error) {
	return handler.GetAllDishes(ctx)
}
