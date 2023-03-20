package handler

import (
	"context"

	"github.com/4klb/coffee-shop-microservices/dish/storage/repo"
	pb "github.com/4klb/coffee-shop-microservices/proto/dish"
)

func GetAllDishes(ctx context.Context) (*pb.DishesResp, error) {
	dishInfo := repo.Dish{}

	allDishes, err := dishInfo.SelectAllDishes(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.DishesResp{
		Dishes: allDishes,
	}, nil
}
