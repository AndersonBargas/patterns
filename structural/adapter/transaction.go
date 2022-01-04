package adapter

type Transaction struct {
	TrackNumber string
	Amount      int
	Status      string
}

func (transaction *Transaction) GetTrackNumber() string {
	return transaction.TrackNumber
}

func (transaction *Transaction) GetAmount() int {
	return transaction.Amount
}

func (transaction *Transaction) GetStatus() string {
	return transaction.Status
}
