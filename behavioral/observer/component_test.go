package observer

import "testing"

func TestDeveCriarUmComponenteReativo(t *testing.T) {
	inputText := NewInputText("country")
	labela := NewLabel("País: {{country}}")
	labelb := NewLabel("Country: {{country}}")
	inputText.Register(labela)
	inputText.Register(labelb)
	inputText.SetValue("Brasil")

	if labela.Value() != "País: Brasil" {
		t.Errorf("Valor deveria ser País: Brasil")
	}

	if labelb.Value() != "Country: Brasil" {
		t.Errorf("Valor deveria ser Country: Brasil")
	}

}
