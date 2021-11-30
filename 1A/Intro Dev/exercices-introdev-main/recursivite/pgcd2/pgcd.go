package pgcd

/*
On peut calculer le pgcd de deux nombres par l'algorithme d'Euclide en utilisant
la propriété suivante : si a > b > 0, pgcd(a, b) = pgcd(b, r) où r est le reste
de la division euclidienne de a par b, et si a > 0, pgcd(a, 0) = a

La fonction pgcd doit implanter cet algorithme simple de manière récursive.

# Entrées
- a : un entier positif ou nul
- b : un entier positif ou nul, tel que a et b ne sont pas tous les deux nuls

# Sorties
- c : le pgcd de a et b
*/

func pgcd(a, b uint) (c uint) {
	if b == 0 {
		return a
	} else {
		var r uint = euclide(a, b)
		if r == 0 {
			return b
		} else {
			a, b = b, r
			c = pgcd(a, b)
		}
	}
	return c
}

func euclide(a, b uint) uint {
	var r uint = a % b
	return r
}
