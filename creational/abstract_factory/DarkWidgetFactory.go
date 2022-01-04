package abstractfactory

type DarkWidgetFactory struct {
}

func NewDarkWidgetFactory() *DarkWidgetFactory {
	return &DarkWidgetFactory{}
}

func (*DarkWidgetFactory) CreateButton() *Button {
	return &Button{
		BackgroundColor: "black",
		Color:           "red",
	}
}

func (*DarkWidgetFactory) CreateLabel() *Label {
	return &Label{
		Color: "red",
	}
}
