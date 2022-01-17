package observer

import "strings"

type Label struct {
	template string
	value    string
}

func NewLabel(template string) *Label {
	return &Label{
		template: template,
	}
}

func (label *Label) Update(key string, value string) {
	label.value = strings.Replace(label.template, "{{"+key+"}}", value, -1)
}

func (label *Label) Value() string {
	return label.value
}
