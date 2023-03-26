package models

//Order ..
type Order struct {
	Id       string  `json:"id"`
	DishId   string  `json:"dishId"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
}
