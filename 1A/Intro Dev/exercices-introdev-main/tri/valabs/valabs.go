package valabs

/*
La fonction valabs doit trier un tableau d'entiers de la plus petite valeur absolue
à la plus grande valeur absolue. En cas
d'égalité de valeur absolue, les nombres
négatifs doivent être placés avant les
nombres positifs.

# Entrée
- tab : un tableau d'entiers
*/

func valabs(tab []int) (trier []int) {

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

func absolue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
