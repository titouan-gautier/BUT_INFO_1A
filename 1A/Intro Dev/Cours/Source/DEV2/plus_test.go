package baseop

import (
	"math/rand"
	"testing"
)

func Test0Plus0(t *testing.T) {
	if Plus(0, 0) != 0 {
		t.Fail()
	}
}

func TestXPlus0(t *testing.T) {
	for x := 1; x < 100; x++ {
		if Plus(x, 0) != x {
			t.Fail()
		}
	}
}

func Test0PlusY(t *testing.T) {
	for y := 1; y < 100; y++ {
		if Plus(0, y) != y {
			t.Fail()
		}
	}
}

func TestXPlusY(t *testing.T) {
	for x := 1; x < 100; x++ {
		for y := 1; y < 100; y++ {
			if Plus(x, y) != x+y {
				t.Fail()
			}
		}
	}
}

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Plus(rand.Intn(100), rand.Intn(100))
	}
}
