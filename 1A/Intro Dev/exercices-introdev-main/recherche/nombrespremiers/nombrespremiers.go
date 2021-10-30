package nombrespremiers

/*
La fonction selectionPremiers filtre le contenu d'un tableau d'entiers pour ne garder que ceux qui sont des nombres premiers.

# Entrée
- t : un tableau d'entiers

# Sortie
- p : un tableau contenant tous les nombres premiers de t, si t est vide, p doit être identique à t

# Exemple
selectionPremiers([]int{1, 2, 3, 4, 5}) = [2 3 5] (l'ordre n'a pas d'importance)
*/

func selectionPremiers(t []int) (p []int) {
	p = make([]int, 0)
	var n bool
	if t == nil {
		return nil
	}
	for i := 0; i < len(t); i++ {
		n = Premier(t[i])
		if n {
			p = append(p, t[i])
		}
	}
	return p
}

func Premier(n int) (prem bool) {
	if n > 1 {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
