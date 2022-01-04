package abstractfactory

type Label struct {
	Color string
}

func NewLabel() *Label {
	return &Label{
		Color: "white",
	}
}
