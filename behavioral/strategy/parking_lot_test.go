package strategy

import (
	"testing"
	"time"
)

func TestDeveCriarUmEstacionamentoVazio(t *testing.T) {
	parkingLot := NewParkingLot(500, NewBeachPriceCalculator())
	if parkingLot.GetEmptySpaces() != 500 {
		t.Error("Devia ter 500 vagas")
	}
}

func TestCarroDeuEntrada(t *testing.T) {
	parkingLot := NewParkingLot(500, NewBeachPriceCalculator())

	dataHoraEntrada, _ := time.Parse("2006-01-02T15:04:05.000Z", "2014-11-12T11:45:26.371Z")
	parkingLot.CheckIn("ABC-12345", dataHoraEntrada)
	if parkingLot.GetEmptySpaces() != 499 {
		t.Error("Devia ter 499 vagas")
	}
}

func TestDeveSairUmCarro(t *testing.T) {
	parkingLot := NewParkingLot(500, NewBeachPriceCalculator())
	parkingLot.CheckIn("ABC-12345", time.Now())
	parkingLot.CheckOut("ABC-12345", time.Now())
	if parkingLot.GetEmptySpaces() != 500 {
		t.Error("Devia ter 500 vagas")
	}
}

func TestDeveCalcularOhValorPorTempoIlimitadoNaPraia(t *testing.T) {
	parkingLot := NewParkingLot(500, NewBeachPriceCalculator())
	parkingLot.CheckIn("ABC-12345", time.Now())
	preco, _ := parkingLot.CheckOut("ABC-12345", time.Now())
	if preco != 20 {
		t.Error("Devia custar 20 reais")
	}
}

func TestDeveCalcularOhValorDeQuatroHorasNoAeroporto(t *testing.T) {
	parkingLot := NewParkingLot(500, NewAirportPriceCalculator())

	dataHoraEntrada := time.Now()
	dataHoraSaida := dataHoraEntrada.Add(time.Hour * 4)

	parkingLot.CheckIn("ABC-12345", dataHoraEntrada)
	preco, _ := parkingLot.CheckOut("ABC-12345", dataHoraSaida)
	if preco != 40 {
		t.Errorf("Devia custar 40 reais, custou %d", preco)
	}
}
