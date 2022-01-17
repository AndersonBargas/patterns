package observer

type Button struct {
}

func NewButton() *Button {
	return &Button{}
}

func (button *Button) Update() {
	println("Button was notified")
}
