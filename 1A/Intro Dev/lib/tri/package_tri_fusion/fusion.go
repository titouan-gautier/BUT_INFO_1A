package fusion

func fusion(tab1 []int, tab2 []int) (trier []int) {
	var i int = 0
	var j int = 0
	var length int
	var lengthfin int

	if len(tab1) > len(tab2) {
		length = len(tab2)
		lengthfin = len(tab1)
	} else {
		length = len(tab1)
		lengthfin = len(tab2)
	}

	for i < length {
		if tab1[i] > tab2[j] {
			trier = append(trier, tab2[j])
			j++
		} else if tab1[i] < tab2[j] {
			trier = append(trier, tab1[i])
			i++
		} else {
			trier = append(trier, tab2[j])
			i++
			j++
		}
	}

	if lengthfin == len(tab2) {
		for j < lengthfin {
			trier = append(trier, tab2[j])
			j++
		}
	} else {
		for i < lengthfin {
			trier = append(trier, tab1[i])
			i++
		}
	}
	return trier
}

func recursif(tab []int) (trier []int) {
	var p int = len(tab) / 2
	for i := 0; i < len(tab); i++ {
		var tab1 []int = tab[:p]
		var tab2 []int = tab[p+1:]
		if len(tab1) != 1 {
			recursif(tab1)
		} else if len(tab2) = 1 {
			fusion(tab1, tab2)
		}
	}

}