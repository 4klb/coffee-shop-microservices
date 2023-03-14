package server

import (
	"context"

	"github.com/4klb/coffeetime/dish/handler"
	"github.com/4klb/coffeetime/proto/dish"
	pb "github.com/4klb/coffeetime/proto/dish"
)

// AddDish implements dish.DishServiceServer
func (s *Server) AddDish(ctx context.Context, in *pb.AddDishReq) (*pb.Empty, error) {
	return &dish.Empty{}, handler.Add(ctx, in)
}

// Dishes implements dish.DishServiceServer
func (s *Server) AllDishes(ctx context.Context, in *pb.Empty) (*pb.DishesResp, error) {
	return handler.GetAllDishes(ctx)
}
