package baseop

import (
	"math/rand"
	"testing"
)

func Test0Mult0(t *testing.T) {
	if Mult(0, 0) != 0 {
		t.Fail()
	}
}

func TestXMult0(t *testing.T) {
	for x := 1; x < 100; x++ {
		if Mult(x, 0) != 0 {
			t.Fail()
		}
	}
}

func Test0MultY(t *testing.T) {
	for y := 1; y < 100; y++ {
		if Mult(0, y) != 0 {
			t.Fail()
		}
	}
}

/*
func TestXMult1(t *testing.T) {
	for x := 0; x < 100; x++ {
		if Mult(x, 1) != x {
			t.Fail()
		}
	}
}

func Test1MultY(t *testing.T) {
	for y := 0; y < 100; y++ {
		if Mult(1, y) != y {
			t.Fail()
		}
	}
}
*/

func BenchmarkMult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mult(rand.Intn(100), rand.Intn(100))
	}
}
