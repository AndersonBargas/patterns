package abstractfactory

type Button struct {
	Color           string
	BackgroundColor string
}

func NewButton() *Button {
	return &Button{
		Color:           "white",
		BackgroundColor: "yellow",
	}
}
