package adapter

type PaypalTransaction struct {
	code        int
	grossAmount int
	situation   string
}

func NewPaypalTransaction(code, grossAmount int, situation string) *PaypalTransaction {
	return &PaypalTransaction{
		code, grossAmount, situation,
	}
}

func (paypalTransaction PaypalTransaction) GetCode() int {
	return paypalTransaction.code
}
