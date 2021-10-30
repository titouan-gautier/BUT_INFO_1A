package nombrespremiers2

/*
La fonction selectionPremiers filtre le contenu d'un tableau d'entiers pour ne garder que ceux qui sont des nombres premiers, en éliminant les doublons.

# Entrée
- t : un tableau d'entiers

# Sortie
- p : un tableau contenant tous les nombres premiers de t, sans doublons.
Si t est vide, p doit être identique à t.

# Exemple
selectionPremiers([]int{1, 2, 2, 3, 4, 5}) = [2 3 5] (l'ordre n'a pas d'importance)
*/

func selectionPremiers(t []int) (p []int) {
	var n bool
	var nb int
	p = make([]int, 0)
	if t == nil {
		return nil
	}
	for i := 0; i < len(t); i++ {
		n = Premier(t[i])
		if n {
			nb = t[i]
			for k := 0; k < len(t); k++ {
				if t[i] == t[k] {
					nb = nb + 1
				}
			}
			if nb == 1 {
				p = append(p, t[i])
			} else {
				for j := 0; j < len(p); j++ {
					if t[i] != p[j] {
						p = append(p, t[i])
					}
				}
			}
		}
	}
	if p == nil {
		p = make([]int, 0)
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
