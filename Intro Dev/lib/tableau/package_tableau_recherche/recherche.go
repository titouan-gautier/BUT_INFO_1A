package recherche

func recherche(t []int, x int) (rep int, res bool) {
	for i := 0; i < len(t); i++ {
		if t[i] == x {
			return i, true
		}
	}
	return -1, false
}
