package adapter

import "fmt"

func NewTransactionFromStripe(stripeTransaction *StripeTransaction) (*Transaction, error) {
	statusMap := make(map[int]string)
	statusMap[1] = "waiting_payhment"
	statusMap[2] = "paid"
	statusMap[3] = "cancelled"

	ourStatus, ok := statusMap[stripeTransaction.situation]
	if !ok {
		return nil, fmt.Errorf("erro ao converter o status: %d", stripeTransaction.situation)
	}

	return &Transaction{
		TrackNumber: stripeTransaction.code,
		Amount:      stripeTransaction.grossAmount,
		Status:      ourStatus,
	}, nil
}
