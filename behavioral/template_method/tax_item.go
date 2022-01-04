package invoice

type TaxItem interface {
	Item
	GetTax() int
}
