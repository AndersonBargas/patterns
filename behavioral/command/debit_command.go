package command

type DebitCommand struct {
	account *Account
	ammount int
}

func NewDebitCommand(account *Account, ammount int) *DebitCommand {
	return &DebitCommand{
		account: account,
		ammount: ammount,
	}
}

func (cc *DebitCommand) Execute() {
	currentBalance := cc.account.Balance()
	cc.account.SetBalance(currentBalance - cc.ammount)
}
