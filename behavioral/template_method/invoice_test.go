package invoice

import "testing"

func TestDeveCriarUmaNotaFiscal(t *testing.T) {
	invoice := NewInvoice()
	invoice.AddItem(NewBeer("Heineken", 10))        // 10% 1
	invoice.AddItem(NewWhisky("Jack Daniels", 100)) // 20% 20
	invoice.AddItem(NewWater("Santa Catarina", 5))  // 0% 0
	taxes := invoice.GetTaxes()
	if taxes != 21 {
		t.Error("Taxas devia ser 21 reais")
	}
}
