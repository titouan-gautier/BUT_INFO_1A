package lib

import (
	"testing"
)

func TestVide(t *testing.T) {
	if !ArbreVide().IsEmpty() {
		t.Error("Test si la liste est bien vide")
	}
}

func Test(t *testing.T) {
	var a = ArbreVide()
	var b = ArbreVide()
	var c = ArbreVide()
	b.Make(2,nil)
	c.Make(3,nil)
	var d []*Arbre = []*Arbre{&b,&c}
	a.Make(1,d)
	if a.valeur != 1 || a.tab[0].valeur != 2 || a.tab[1].valeur != 3 {
		t.Error("Fail")
	}
}