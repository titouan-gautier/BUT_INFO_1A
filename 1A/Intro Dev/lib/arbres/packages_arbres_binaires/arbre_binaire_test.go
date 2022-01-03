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
	b.Make(2,nil,nil)
	c.Make(3,nil,nil)
	a.Make(1,&b,&c)
	if a.val != 1 || a.gauche.val != 2 || a.droite.val != 3 {
		t.Error("Fail")
	}
}