package insertionenplace

import "testing"

func Test1(t *testing.T) {
	var tab []int = []int{5, 4, 3, 2, 1}
	var tri []int = insertion(tab)
	if verif(tri, []int{1, 2, 3, 4, 5}) != true {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	var tab []int = []int{5, 4, 3, 2, 1}
	var tri []int = insertion(tab)
	if verif(tri, []int{2, 1, 3, 4, 5}) != false {
		t.Fail()
	}
}

func TestVide1(t *testing.T) {
	var tab []int
	var tri []int = insertion(tab)
	if verif(tri, []int{2, 1, 3, 4, 5}) != false {
		t.Fail()
	}
}