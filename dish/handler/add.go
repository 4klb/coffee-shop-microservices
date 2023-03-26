package handler

import (
	"context"
	"encoding/json"
	"log"

	uuid "github.com/satori/go.uuid"

	pb "github.com/4klb/coffee-shop-microservices/proto/dish"

	"github.com/4klb/coffee-shop-microservices/dish/broker"
	"github.com/4klb/coffee-shop-microservices/dish/config"
	"github.com/4klb/coffee-shop-microservices/dish/storage/repo"
)

func Add(ctx context.Context, req *pb.AddDishReq) error {
	//todo req check

	dishInfo := repo.Dish{
		Id:          uuid.Must(uuid.NewV4(), nil).String(),
		Name:        req.Dish.Name,
		Description: req.Dish.Description,
		Price:       req.Dish.Price,
	}

	// if err := dishInfo.InsertDish(ctx); err != nil {
	// 	return err
	// }

	rmq, err := broker.GetRabbitConnection()
	if err != nil {
		log.Fatalf("broker.GetRabbitConnection: %s", err)
	}

	publisher, err := broker.NewPublisher(rmq)
	if err != nil {
		log.Fatalf("broker.NewPublisher: %s", err)
	}

	msg, err := json.Marshal(dishInfo)
	if err != nil {
		return err
	}

	publisher.Created(context.Background(), config.GetConfig().RabbitMQ.RoutingKey, msg)

	return nil
}
