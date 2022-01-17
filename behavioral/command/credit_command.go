package command

type CreditCommand struct {
	account *Account
	ammount int
}

func NewCreditCommand(account *Account, ammount int) *CreditCommand {
	return &CreditCommand{
		account: account,
		ammount: ammount,
	}
}

func (cc *CreditCommand) Execute() {
	currentBalance := cc.account.Balance()
	cc.account.SetBalance(currentBalance + cc.ammount)
}
