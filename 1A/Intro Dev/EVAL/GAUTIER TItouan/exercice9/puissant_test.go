package puissant

import "testing"

func TestVide(t *testing.T) {
	if !puissant("") {
		t.Error("Le mot vide est puissant (pour toutes les puissances)")
	}
}

func TestCarre(t *testing.T) {
	if !puissant("bonbon") {
		t.Error("Le mot \"bonbon\" est puissant (puissance 2)")
	}
}

func TestCube(t *testing.T) {
	if !puissant("lalala") {
		t.Error("Le mot \"lalala\" est puissant (puissance 3)")
	}
}

func TestFaux(t *testing.T) {
	if puissant("lalalilala") {
		t.Error("Le mot \"lalalilala\" n'est pas puissant")
	}
}
