package doublons

import "testing"

func TestVide(t *testing.T) {
	if !doublons(nil) {
		t.Error("doublons(nil) devrait retourner true")
	}
	if !doublons([]int{}) {
		t.Error("doublons([]int{}) devrait retourner true")
	}
}

func TestDouble(t *testing.T) {
	if doublons([]int{1, 1}) {
		t.Error("doublons([]int{1, 1}) devrait retourner false")
	}
	if doublons([]int{1, 2, 3, 4, 5, 1}) {
		t.Error("doublons([]int{1, 2, 3, 4, 5, 1}) devrait retourner false")
	}
}

func TestOk(t *testing.T) {
	if !doublons([]int{1, 2, 3, 4, 5, 6}) {
		t.Error("doublons([]int{1, 2, 3, 4, 5, 6}) devrait retourner true")
	}
}

func TestDehors(t *testing.T) {
	if doublons([]int{2, 3, 4, 5, 6, 7}) {
		t.Error("doublons([]int{2, 3, 4, 5, 6, 7}) devrait retourner false")
	}
}
