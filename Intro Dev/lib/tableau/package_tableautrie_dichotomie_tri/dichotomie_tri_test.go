package dichotomie_tri

import (
	"testing"
)

func Test1(t *testing.T) {
	var tab []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var n = 3
	var b bool
	var a int
	b, a = dichotomie_tri(tab, n)
	if b != true || a != 2 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	var tab []int = []int{10, 20, 30, 100, 500}
	var n = 100
	var b bool
	var a int
	b, a = dichotomie_tri(tab, n)
	if b != true || a != 3 {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	var tab []int = []int{10, 20, 30, 100, 500}
	var n = 1000
	var b bool
	var a int
	b, a = dichotomie_tri(tab, n)
	if b != false || a != -1 {
		t.Fail()
	}
}

func TestVide1(t *testing.T) {
	var tab []int
	var n = 100
	var b bool
	var a int
	b, a = dichotomie_tri(tab, n)
	if b != false || a != -1 {
		t.Fail()
	}
}

func TestVide2(t *testing.T) {
	var tab []int = make([]int, 0)
	var n = 100
	var b bool
	var a int
	b, a = dichotomie_tri(tab, n)
	if b != false || a != -1 {
		t.Fail()
	}
}
