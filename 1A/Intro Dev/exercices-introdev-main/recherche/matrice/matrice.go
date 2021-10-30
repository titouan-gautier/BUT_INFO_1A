package matrice

import (
	"errors"
)

var errPasMat error = errors.New("mat n'est pas vraiment une matrice")

/*
La fonction compte indiquera combien de fois un nombre n apparaît dans une matrice.

# Entrées
- n : le nombre à chercher
- mat : une matrice d'entiers (un tableau de tableaux qui font tous la même taille)

# Sorties
- num : le nombre de fois que n apparaît dans mat
- err : errPasMat si mat n'est pas une vraie matrice (toutes les lignes n'ont pas
la même longueur ou mat vaut nil ou une ligne vaut nil)

# Exemple
compte(0, [][]int{[]int{0, 1}, []int{0, 0}}) = 3
*/
/*func compte(n int, mat [][]int) (num int, err error) {
	var i int
	//var k int
	if mat != nil {
		for i = 0; i < len(mat[i]); i++ {
			if mat[i] != nil {
				if len(mat[i+1]) != len(mat[i]) {
					return 0, errPasMat
				}
			} else {
				return 0, errPasMat
			}

		}
	} else {
		return 0, errPasMat
	}

	return num, err
}
*/
func compte(n int, tab [][]int) (num int, err error) {
	if tab == nil {
		return 0, errPasMat
	}
	for i := 0; i < len(tab); i++ {
		if tab[i] == nil {
			return 0, errPasMat
		}
		for k := 0; k < len(tab[i]); k++ {
			if tab[i][k] == n {
				num = num + 1
			}
		}
	}

	return num, err
}
