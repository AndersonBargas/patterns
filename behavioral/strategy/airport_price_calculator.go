package strategy

type AirportPriceCalculator struct {
}

func NewAirportPriceCalculator() *AirportPriceCalculator {
	return &AirportPriceCalculator{}
}

func (calculator *AirportPriceCalculator) Calculate(hours int) int {
	return hours * 10
}
