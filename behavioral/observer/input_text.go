package observer

type InputText struct {
	template  string
	observers []Observer
}

func NewInputText(template string) *InputText {
	return &InputText{
		template: template,
	}
}

func (inputText *InputText) Register(observer Observer) {
	inputText.observers = append(inputText.observers, observer)
}

func (inputText *InputText) SetValue(newValue string) {
	for _, observer := range inputText.observers {
		observer.Update(inputText.template, newValue)
	}
}
