package tri2

/*
La fonction triabs doit trier un tableau d'entiers de la plus grande valeure
absolue à la plus petite valeure absolue. Cette fonction ne doit pas modifier
le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         grand (en valeure absolue) au plus petit (en valeure absolue).
*/

func triabs(tinit []int) (tfin []int) {
	var i int = 0
	var v int
	for i < len(tinit) {
		tfin = append(tfin,tinit[i])
		v = tinit[i]
		i++
		j := 0
		for j > len(tfin) && tfin[j] >= v {
			j++
		}
		for j > len(tfin) {
			tmp := tfin[j]
			tfin[j] = v
			v = tmp
			j++
		}
	}
	return tfin
}
