package nombrespairs

/*
La fonction sommeNombresPairs doit retourner la somme de tous les nombres
pairs contenus entre deux entiers (inclus).

# Entrées
- a : un nombre entier, une des bornes
- b : un nombre entier, l'autre borne

# Sortie
- somme : la somme de tous les nombres entiers compris entre a et b inclus

# Exemple
sommeNombresPairs(2, 7) = 12
*/
func sommeNombresPairs(a, b int) (somme int) {
	var aa int = a
	var bb int = b
	if a > b {
		aa = b
		bb = a
	}
	for i := aa; i <= bb; i++ {
		if i%2 == 0 {
			somme = somme + i
		}
	}
	return somme
}
