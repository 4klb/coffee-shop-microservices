package handler

import (
	"context"

	uuid "github.com/satori/go.uuid"

	pb "github.com/4klb/coffee-shop-microservices/proto/dish"

	"github.com/4klb/coffee-shop-microservices/dish/storage/repo"
)

func Add(ctx context.Context, req *pb.AddDishReq) error {
	//todo validate

	dishInfo := repo.Dish{
		Id:          uuid.Must(uuid.NewV4(), nil).String(),
		Name:        req.Dish.Name,
		Description: req.Dish.Description,
		Price:       req.Dish.Price,
	}

	if err := dishInfo.InsertDish(ctx); err != nil {
		return err
	}

	return nil
}
