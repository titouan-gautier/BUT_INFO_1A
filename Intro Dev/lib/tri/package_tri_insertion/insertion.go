package insertion

func insertion(tab []int) (trier []int) {
	
	if tab == nil {
		return nil
	}
	if len(tab) == 0 {
		return nil
	}

	for i := 0; i < len(tab); i++ {
		trier = append(trier, tab[i])
		var j int = 0
		var v int = tab[i]
		for j < len(trier) && trier[j] <= v {
			j++
		}
		for j < len(trier) {
			tmp := trier[j]
			trier[j] = v
			v = tmp
			j++
		}
	}
	return trier
}
