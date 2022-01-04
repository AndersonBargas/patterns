package abstractfactory

type LightWidgetFactory struct {
}

func NewLightWidgetFactory() *LightWidgetFactory {
	return &LightWidgetFactory{}
}

func (*LightWidgetFactory) CreateButton() *Button {
	return &Button{
		BackgroundColor: "white",
		Color:           "black",
	}
}

func (*LightWidgetFactory) CreateLabel() *Label {
	return &Label{
		Color: "black",
	}
}
