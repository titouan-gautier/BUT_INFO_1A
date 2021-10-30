package insertion

import "testing"

//go test -coverprofile=couverture.out
//go tool cover -html=couverture.out

func Test1(t *testing.T) {
	var tab []int = []int{4, 66, 44, 12}
	var trier []int = insertion(tab)
	for verif(trier, []int{4, 12, 44, 66}) != true {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	var tab []int = []int{4, 66, 44, 12}
	var trier []int = insertion(tab)
	if verif(trier, []int{4, 12, 44, 66}) != true {
		if verif(tab,[]int{4,66,44,12}) {
			t.Fail()
		} 
	}
}
func TestVide1(t *testing.T) {
	var tab []int
	var trier []int = insertion(tab)
	for verif(trier, nil) != true {
		t.Fail()
	}
}

func TestVide2(t *testing.T) {
	var tab []int = []int{}
	var trier []int = insertion(tab)
	for verif(trier, nil) != true {
		t.Fail()
	}
}

func verif(tab1 []int, tab2 []int) (b bool) {
	if len(tab1) != len(tab2) {
		return false
	}
	for i := 0; i < len(tab1); i++ {
		if tab1[i] != tab2[i] {
			return false
		}
	}
	return true
}
