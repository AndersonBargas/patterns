package adapter

type StripeTransaction struct {
	code        string
	grossAmount int
	situation   int
}

func NewStripeTransaction(code string, grossAmount, situation int) *StripeTransaction {
	return &StripeTransaction{
		code, grossAmount, situation,
	}
}

func (stripeTransaction StripeTransaction) GetCode() string {
	return stripeTransaction.code
}
