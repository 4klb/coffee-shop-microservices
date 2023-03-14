package repo

import (
	"context"
	"log"

	"github.com/4klb/coffeetime/dish/storage/postgres"

	pb "github.com/4klb/coffeetime/proto/dish"
)

type Dish struct {
	Id          string
	Name        string
	Description string
	Price       float32
}

func (d *Dish) InsertDish(ctx context.Context) error {
	db := postgres.GetDB(ctx)

	_, err := db.ExecContext(ctx, `
	INSERT INTO dish (dish_id, dish_name, dish_description, dish_price) 
	VALUES ($1, $2, $3, $4);`, d.Id, d.Name, d.Description, d.Price)
	if err == context.DeadlineExceeded {
		return context.DeadlineExceeded
	}

	if err != nil {
		return err
	}
	return nil
}

func (d *Dish) SelectAllDishes(ctx context.Context) ([]*pb.Dish, error) {
	var dishesProto []*pb.Dish

	db := postgres.GetDB(ctx)

	rows, err := db.QueryContext(ctx, `SELECT * FROM dish;`)
	if err == context.DeadlineExceeded {
		return nil, context.DeadlineExceeded
	}
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var dishProto *pb.Dish
		err := rows.Scan(&dishProto.Id, &dishProto.Name, &dishProto.Description, &dishProto.Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		dishesProto = append(dishesProto, dishProto)
	}

	return dishesProto, nil
}
