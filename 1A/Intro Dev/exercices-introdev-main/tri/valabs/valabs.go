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

func valabs(t []int) (tri []int) {
	var i int = 1
	for i < len(t) {
		v := t[i]
		j := i - 1
		for j >= 0 && t[j] > absolue(v) {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = v
		i++
	}
	return t
}

func absolue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
