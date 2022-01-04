package invoice

type Invoice struct {
	items []Item
}

func NewInvoice() *Invoice {
	return &Invoice{}
}

func (invoice *Invoice) AddItem(item Item) {
	invoice.items = append(invoice.items, item)
}

func (invoice *Invoice) GetTaxes() int {
	taxes := 0
	for _, item := range invoice.items {
		if convertedItem, isTaxed := item.(TaxItem); isTaxed {
			price := convertedItem.GetPrice()
			tax := convertedItem.GetTax()
			taxes += (price * tax) / 100
		}
	}
	return taxes
}
