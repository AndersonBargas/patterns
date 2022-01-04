package invoice

type Water struct {
	description string
	price       int
}

func NewWater(description string, price int) *Water {
	return &Water{description, price}
}

func (water *Water) GetPrice() int {
	return water.price
}
