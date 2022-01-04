package invoice

type Whisky struct {
	description string
	price       int
}

func NewWhisky(description string, price int) *Whisky {
	return &Whisky{description, price}
}

func (whisky *Whisky) GetPrice() int {
	return whisky.price
}

func (whisky *Whisky) GetTax() int {
	return 20
}
