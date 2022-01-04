package strategy

type PriceCalculatorInterface interface {
	Calculate(hours int) int
}
