package selection

import "testing"

func Test1(t *testing.T) {
	var tab []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var n int = 5
	var tb []int
	var b bool
	tb, b = selection(tab, n)
	if b != true || verif(tb, []int{1, 2, 3, 4, 5}) != true {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	var tab []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var n int = 5
	var tb []int
	var b bool
	tb, b = selection(tab, n)
	if b != true || verif(tb, []int{2, 2, 3, 4, 5}) != false {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	var tab []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var n int = 44
	var tb []int
	var b bool
	tb, b = selection(tab, n)
	if b != false || verif(tb, []int{1, 2, 3, 4, 5, 6, 7, 8}) != true {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	var tab []int = []int{10, 11, 12, 13}
	var n int = 9
	var tb []int
	var b bool
	tb, b = selection(tab, n)
	if b != false || verif(tb, []int{}) != true {
		t.Fail()
	}
}

func TestVide1(t *testing.T) {
	var tab []int
	var n int = 5
	var tb []int
	var b bool
	tb, b = selection(tab, n)
	if b != false || tb != nil {
		t.Fail()
	}
}

func TestVide2(t *testing.T) {
	var tab []int = make([]int, 0)
	var n int = 5
	var tb []int
	var b bool
	tb, b = selection(tab, n)
	if b != false || tb != nil {
		t.Fail()
	}
}
