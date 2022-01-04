package adapter

import "testing"

func TestDeveCriarUmaTransacaoDoStripe(t *testing.T) {
	stripeTransaction := NewStripeTransaction("AHN786AB8", 1000, 2)
	if stripeTransaction.GetCode() != "AHN786AB8" {
		t.Error("Erro ao criar a transacao no Stripe")
	}
}

func TestDeveCriarUmaTransacaoDoPaypal(t *testing.T) {
	paypalTransaction := NewPaypalTransaction(78978978, 1000, "S")
	if paypalTransaction.GetCode() != 78978978 {
		t.Error("Erro ao criar a transacao no Paypal")
	}
}

func TestDeveCriarUmaTransacaoAhPartirDoStripe(t *testing.T) {
	stripeTransaction := NewStripeTransaction("AHN786AB8", 1000, 2)
	transaction, err := NewTransactionFromStripe(stripeTransaction)
	if err != nil {
		t.Error("Nao devia ter erro na conversao")
	}

	if transaction.GetTrackNumber() != "AHN786AB8" {
		t.Error("Codigo da transacao nao bateu")
	}

	if transaction.GetAmount() != 1000 {
		t.Error("Valor da transacao nao bateu")
	}

	if transaction.GetStatus() != "paid" {
		t.Error("Situacao da transacao nao bateu")
	}
}
