package insertionenplace

func insertion(t []int) (tri []int) {
	var i int = 1
	for i < len(t) {
		v := t[i]
		j := i - 1
		for j >= 0 && t[j] > v {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = v
		i++
	}
	return t
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