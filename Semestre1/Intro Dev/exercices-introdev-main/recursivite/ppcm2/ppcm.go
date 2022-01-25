package ppcm

/*
On peut calculer le ppcm de deux nombres x0 et y0, en appliquant la méthode
récursive suivante :
- ppcm(x, y) = x si x et y sont égaux
- ppcm(x, y) = ppcm(x + x0, y) si x < y
- ppcm(x, y) = ppcm(x, y + y0) si x > y
- on commence par appeler ppcm(x0, y0)
Cette méthode se généralise au calcul du ppcm de n nombres.

La fonction ppcm est une fonction récursive qui calcule le ppcm de n nombres à
partir de la technique décrite ci-dessus.

Vous pouvez considérer que la fonction ppcm ne sera jamais appelée avec un des
nombres considérés qui soit égal à 0.

# Entrées
- tab0 : un tableau d'entiers (ce tableau ne sera jamais vide)
PGCD
# Sortie
- z : le ppcm des entiers contenus dans tab0
*/

func real(tab0, tab []uint) (z uint) {
	var i int = 0
	if len(tab) == 2 {
		return z
	} else {
		x := tab[i]
		y := tab[i+1]
		if x < y {
			tab[i] = tab0[i] + tab[i]
			z = real(tab0, tab)
		}
		if x > y {
			tab[i+1] = tab0[i+1] + tab[i+1]
			z = real(tab0, tab)
		}
	}
	return z
}

func ppcm(tab0 []uint) (z uint) {
	var tab []uint = tab0
	z = real(tab0, tab)
	return z
}
