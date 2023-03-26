package models

//Order ..
type Order struct {
	Id       string
	DishID   string
	Quantity int
	Price    float32
	Discount float32
}
