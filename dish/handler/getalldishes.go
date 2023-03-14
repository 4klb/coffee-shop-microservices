package handler

import (
	"context"

	"github.com/4klb/coffeetime/dish/storage/repo"
	pb "github.com/4klb/coffeetime/proto/dish"
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
