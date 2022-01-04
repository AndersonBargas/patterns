package abstractfactory

import "testing"

func TestViewComTemaClaro(t *testing.T) {
	view := NewView("light")

	if view.Button.Color != "black" {
		t.Error("Botao devia ter o texto na cor preta")
	}

	if view.Button.BackgroundColor != "white" {
		t.Error("Botao devia ter o fundo na cor branca")
	}

	if view.Label.Color != "black" {
		t.Error("Label devia ter o texto na cor preta")
	}
}

func TestViewComTemaEscuro(t *testing.T) {
	view := NewView("dark")

	if view.Button.Color != "red" {
		t.Error("Botao devia ter o texto na cor preta")
	}

	if view.Button.BackgroundColor != "black" {
		t.Error("Botao devia ter o fundo na cor branca")
	}

	if view.Label.Color != "red" {
		t.Error("Label devia ter o texto na cor preta")
	}
}
