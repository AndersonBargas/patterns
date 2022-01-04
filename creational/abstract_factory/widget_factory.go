package abstractfactory

type WidgetFactoryInterface interface {
	CreateButton() *Button
	CreateLabel() *Label
}

type WidgetFactory struct {
	Button *Button
	Label  *Label
}

func NewWidgetFactory(widgetFactoryInterface WidgetFactoryInterface) *WidgetFactory {
	return &WidgetFactory{
		Button: widgetFactoryInterface.CreateButton(),
		Label:  widgetFactoryInterface.CreateLabel(),
	}
}

func (widgetFactory *WidgetFactory) GetButton() *Button {
	return widgetFactory.Button
}

func (widgetFactory *WidgetFactory) GetLabel() *Label {
	return widgetFactory.Label
}
