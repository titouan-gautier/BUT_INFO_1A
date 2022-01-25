package decroissant

/*
La fonction décroissant doit trier un tableau d'entiers du plus grand
au plus petit.

# Entrée
- tab : le tableau à trier
*/

func decroissant(tab []int) (trier []int) {
	var i int = 1
	for i < len(tab) {
		v := tab[i]
		j := i - 1
		for j >= 0 && tab[j] < v {
			tab[j+1] = tab[j]
			j--
		}
		tab[j+1] = v
		i++
	}
	return trier
}
