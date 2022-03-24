package nombrespremiers

/*
La fonction premiers doit retourner un tableau contenant les nombres
premiers compris entre 0 et un entier n.

# Entrée
- n : un nombre entier

# Sortie
- p : un tableau contenant tous les nombres premiers compris entre 0 et n inclus, s'il n'existe pas de tels nombres premiers, il faut que p soit un tableau contenant 0 éléments

# Exemple
premiers(10) = [2 3 5 7] (l'ordre n'a pas d'importance)
*/
func premiers(n int) (p []int) {
	if n >= 2 {
		for i := 0; i <= n; i++ {
			if estPremier(i) {
				p = append(p, i)
			}
		}
		return p
	} else {
		p = make([]int, 0)
		return p
	}

}

func estPremier(n int) (b bool) {
	var nb int
	var comp int = n + 1
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			nb = nb + i
		}
	}
	if nb > 0 {
		if nb == comp {
			b = true
		} else {
			b = false
		}
	} else {
		b = false
	}
	return b
}
