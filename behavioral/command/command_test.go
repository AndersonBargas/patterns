package command

import "testing"

func TestDeveCriarUmaConta(t *testing.T) {
	account := NewAccount(0)
	if account.Balance() != 0 {
		t.Errorf("Conta devia estar zerada")
	}
}

func TestDeveDebitarUmaConta(t *testing.T) {
	account := NewAccount(100)
	account.SetBalance(50)
	if account.Balance() != 50 {
		t.Errorf("Devia ter sobrado 50")
	}
}

func TestDeveCreditarUmaContaUsandoUmComando(t *testing.T) {
	account := NewAccount(0)
	creditCommand := NewCreditCommand(account, 100)
	creditCommand.Execute()
	debitCommand := NewDebitCommand(account, 50)
	debitCommand.Execute()
	if account.Balance() != 50 {
		t.Errorf("Devia ter sobrado 50")
	}
}
