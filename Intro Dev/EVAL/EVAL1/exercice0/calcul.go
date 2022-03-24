package calcul

/*
La fonction calcul doit retourner la valeur de son argument multipliée par 3, à laquelle on ajoute ensuite 5.

# Entrée
- x : un entier

# Sortie
- res : 3 fois x plus 5
*/

func calcul(x int) (res int) {
	res = (x*3) + 5
	return res
}
