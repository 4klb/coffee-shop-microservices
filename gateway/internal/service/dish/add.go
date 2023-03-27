package dish

import (
	"context"

	"github.com/4klb/coffee-shop-microservices/dish/config"
	"github.com/4klb/coffee-shop-microservices/gateway/internal/service"
	"github.com/4klb/coffee-shop-microservices/proto/dish"
	"github.com/4klb/coffee-shop-microservices/utils"
)

func Add(ctx context.Context, req *dish.AddDishReq) error {
	addr := config.GetConfig().DishServer.Port

	conn, err := service.GetConnection(ctx, addr)
	defer utils.MuteCloseClientConn(conn)

	if err != nil {
		return err
	}

	client := dish.NewDishServiceClient(conn)

	_, err = client.AddDish(ctx, req)

	return err
}
