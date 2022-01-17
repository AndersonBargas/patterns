package observer

type Expression struct {
}

func NewExpression() *Expression {
	return &Expression{}
}

func (expression *Expression) Update() {
	println("Expression was notified")
}
