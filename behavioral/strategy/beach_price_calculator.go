package strategy

type BeachPriceCalculator struct {
}

func NewBeachPriceCalculator() *BeachPriceCalculator {
	return &BeachPriceCalculator{}
}

func (calculator *BeachPriceCalculator) Calculate(hours int) int {
	return 20
}
