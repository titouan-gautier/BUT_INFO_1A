package minimum

import "testing"

func Test2(t *testing.T) {
	var tab []int = []int{73, 23, 65, 12, 3}
	var a int
	var b bool
	b, a = minimum(tab)
	if a != 3 || b != true {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	var tab []int = []int{-4, -2, -8, -1}
	var a int
	var b bool
	b, a = minimum(tab)
	if a != -8 || b != true {
		t.Fail()
	}
}

func Vide1(t *testing.T) {
	var tab []int = []int{}
	var a int
	var b bool
	b, a = minimum(tab)
	if a != 0 || b != false {
		t.Fail()
	}
}

func Vide2(t *testing.T) {
	var tab []int
	var a int
	var b bool
	b, a = minimum(tab)
	if b != false || a != -1 {
		t.Fail()
	}
}