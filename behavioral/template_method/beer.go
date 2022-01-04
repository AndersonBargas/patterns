package invoice

type Beer struct {
	description string
	price       int
}

func NewBeer(description string, price int) *Beer {
	return &Beer{description, price}
}

func (beer *Beer) GetPrice() int {
	return beer.price
}

func (beer *Beer) GetTax() int {
	return 10
}
