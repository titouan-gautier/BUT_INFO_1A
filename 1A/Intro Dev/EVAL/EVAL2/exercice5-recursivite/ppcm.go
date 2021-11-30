package ppcm

/*
On peut calculer le ppcm de deux nombres x0 et y0, en appliquant la méthode
récursive suivante :
- ppcm(x, y) = x si x et y sont égaux
- ppcm(x, y) = ppcm(x + x0, y) si x < y
- ppcm(x, y) = ppcm(x, y + y0) si x > y
- on commence par appeler ppcm(x0, y0)

La fonction ppcm doit être une fonction récursive (ou une fonction qui se base
sur une fonction récursive) qui calcule le ppcm de deux nombres à partir de la
technique décrite ci-dessus. Une solution qui n'est pas récursive ne rapportera
pas de points.

Vous pouvez considérer que la fonction ppcm ne sera jamais appelée avec un de ses
arguments qui soit égal à 0.

# Entrées
- x : un entier
- y : un entier

# Sortie
- z : le ppcm de x et y
*/

func ppcm(x, y uint) (z uint) {
	var x0 uint = x
	var y0 uint = y
	return oui(x, y, x0, y0)
}

func oui(x, y, x0, y0 uint) (z uint) {
	if x == y {
		return x
	} else {
		if x < y {
			z = oui(x+x0, y, x0, y0)
		} else {
			z = oui(x, y+y0, x0, y0)
		}
	}
	return z
}
