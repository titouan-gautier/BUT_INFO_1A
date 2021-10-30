package rechercheplus

import "testing"

func Test1(t *testing.T) {
	var tab []int = []int{0, 2, 4, 1, 3}
	var a int
	var b bool
	a, b = rechercheplus(tab, 2)
	if b != true || a != 1 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	var tab []int
	var a int
	var b bool
	a, b = rechercheplus(tab, 1)
	if b != false || a != -1 {
		t.Fail()
	}

}

func Test3(t *testing.T) {
	var tab []int = []int{0, 2, 4, 1, 3}
	var a int
	var b bool
	a, b = rechercheplus(tab, 12)
	if b != false || a != -1 {
		t.Fail()
	}
}
