package command

type Account struct {
	balance int
}

func NewAccount(startBalance int) *Account {
	return &Account{
		balance: startBalance,
	}
}

func (account *Account) Balance() int {
	return account.balance
}

func (account *Account) SetBalance(newBalance int) {
	account.balance = newBalance
}
