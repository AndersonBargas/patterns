package strategy

import (
	"fmt"
	"time"
)

type ParkingLot struct {
	calculator PriceCalculatorInterface
	capacity   int
	// [placa]horaDoCheckin
	checkins map[string]time.Time
}

func NewParkingLot(capacity int, calculator PriceCalculatorInterface) *ParkingLot {
	return &ParkingLot{
		calculator: calculator,
		capacity:   capacity,
		checkins:   make(map[string]time.Time),
	}
}

func (parkingLot *ParkingLot) CheckIn(placa string, dataHora time.Time) {
	parkingLot.checkins[placa] = dataHora
}

func (parkingLot *ParkingLot) CheckOut(placa string, dataHora time.Time) (int, error) {
	if _, ok := parkingLot.checkins[placa]; !ok {
		return 0, fmt.Errorf("o carro de placa %s nao foi encontrado", placa)
	}

	diferenca := dataHora.Sub(parkingLot.checkins[placa])
	delete(parkingLot.checkins, placa)
	return parkingLot.calculator.Calculate(int(diferenca.Hours())), nil
}

func (parkingLot *ParkingLot) GetEmptySpaces() int {
	return parkingLot.capacity - len(parkingLot.checkins)
}
