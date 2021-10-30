package dichotomie_tri

func dichotomie_tri(tab []int, x int) (b bool, n int) {
	var start int = 0
	var end int = len(tab)

	if tab == nil {
		return false, -1
	}

	if len(tab) == 0 {
		return false, -1
	}

	for start < end {
		milieu := (start + end) / 2
		if tab[milieu] == x {
			return true, milieu
		} else if tab[milieu] > x {
			end = milieu
		} else if tab[milieu] < x {
			start = milieu + 1
		}
	}
	return false, -1
}
