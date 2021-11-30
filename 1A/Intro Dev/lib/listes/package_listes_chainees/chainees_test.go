package lib

import "testing"
import "log"

func Test1(t *testing.T) {
	if !ListeVide().VerifVide() {
		t.Error("Test si la liste est bien vide avec la fonction VerifListe")
	}
}

func Test2(t *testing.T) {
	var l Liste_chainée = ListeVide()
	var s string = "10 9 8 7 6 5 4 3 2 1 0"

	for i := 0; i <= 10; i++ {
		l = Append(i, l)
	}

	log.Print(l.Tail().Tail())

	lfin := l.Affiche()

	if lfin != s {
		t.Error("Test la fonction Append")
	}
}

func Test3(t *testing.T) {
	var l Liste_chainée = ListeVide()
	var s int = 10

	for i := 0; i <= 10; i++ {
		l = Append(i, l)
	}
	if l.Head() != s {
		t.Error("Test la fonction Head")
	}
}
