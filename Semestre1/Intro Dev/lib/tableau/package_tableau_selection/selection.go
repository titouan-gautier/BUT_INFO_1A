package selection

/* Cette fonction va retourner le tableau de tout les nombre plus petit que x compris dans notre tableau
 */

func selection(tab []int, n int) (t []int, b bool) {
	if tab == nil {
		return nil, false
	}
	
	if len(tab) == 0 {
		return nil, false
	}

	for i := 0; i < len(tab); i++ {
		if tab[i] <= n {
			t = append(t, tab[i])
		}
	}
	if verif(tab, t) {
		return t, false
	}
	if len(t) == 0 {
		return t, false
	}
	return t, true
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
