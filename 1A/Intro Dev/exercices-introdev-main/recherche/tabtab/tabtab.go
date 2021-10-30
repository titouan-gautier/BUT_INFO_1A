package tabtab

import (
	"errors"
)

var errPasTabTab error = errors.New("tab n'est pas vraiment un tableau de tableaux")

/*
La fonction compte indiquera combien de fois un nombre n apparaît dans une tableau
de tableaux.

# Entrées
- n : le nombre à chercher
- tab : un tableau de tableau

# Sorties
- num : le nombre de fois que n apparaît dans tab
- err : errPasTabTab si tab vaut nil ou si une de ses lignes vaut nil

# Exemple
compte(0, [][]int{[]int{0, 1}, []int{0, 0, 0}}) = 4
*/
func compte(n int, tab [][]int) (num int, err error) {
	if tab == nil {
		return 0, errPasTabTab
	}
	for i := 0; i < len(tab); i++ {
		if tab[i] == nil {
			return 0, errPasTabTab
		}
		for k := 0; k < len(tab[i]); k++ {
			if tab[i][k] == n {
				num = num + 1
			}
		}
	}

	return num, err
}

