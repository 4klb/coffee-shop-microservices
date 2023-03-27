package api

import (
	"context"
	"net/http"
	"time"

	"github.com/4klb/coffee-shop-microservices/gateway/internal/models"
	"github.com/4klb/coffee-shop-microservices/gateway/internal/response"
	"github.com/gin-gonic/gin"

	"github.com/4klb/coffee-shop-microservices/gateway/internal/service/dish"
	pb "github.com/4klb/coffee-shop-microservices/proto/dish"
)

const defaultCtx = time.Second * 10

//AddDish ..
func (h *Handle) AddDish(c *gin.Context) {
	var dishreq models.Dish

	if err := c.BindJSON(&dishreq); err != nil {
		response.ResponseWithError(http.StatusBadRequest, err, c)
		return
	}

	if err := HandleAddDish(dishreq); err != nil {
		response.ResponseWithError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data was added"})
}

func HandleAddDish(dishreq models.Dish) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtx)
	defer cancel()

	dishProto := &pb.Dish{
		Id:          dishreq.Id,
		Name:        dishreq.Name,
		Description: dishreq.Description,
		Price:       dishreq.Price,
	}

	if err := dish.Add(ctx, &pb.AddDishReq{
		Dish: dishProto,
	}); err != nil {
		return err
	}

	return nil
}
