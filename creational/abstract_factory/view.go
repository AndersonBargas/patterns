package abstractfactory

type View struct {
	Button *Button
	Label  *Label
}

func NewView(theme string) *View {
	if theme == "light" {
		widgetFactory := NewWidgetFactory(NewLightWidgetFactory())
		return &View{
			Button: widgetFactory.GetButton(),
			Label:  widgetFactory.GetLabel(),
		}
	}

	if theme == "dark" {
		widgetFactory := NewWidgetFactory(NewDarkWidgetFactory())
		return &View{
			Button: widgetFactory.GetButton(),
			Label:  widgetFactory.GetLabel(),
		}
	}

	return nil

}
