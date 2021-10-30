package minimum

/* Cette fonction va trouver le plus petit nombre dans un tableau.
 */

func minimum(tab []int) (b bool, min int) {
	min = tab[0]
	for i := 1; i < len(tab); i++ {
		if tab[i] < min {
			min = tab[i]
		}
	}
	return true, min
}
