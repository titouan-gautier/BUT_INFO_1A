package bienrange

/*
La fonction bienrange indique si un tableau d'entiers est bien trié par ordre
croissant ou pas.

# Entrée
- tab : le tableau d'entiers à analyser

# Sortie
- estrange : un booléen qui vaut true si les entiers de tab sont bien triés en
ordre croissant et false s'ils ne sont pas bien triés.
*/

func bienrange(tab []int) (estrange bool) {
	var croissant bool = true
	var decroissant bool = true

	for i := 1; i < len(tab) && croissant; i++ {
		croissant = tab[i-1] <= tab[i]
	}
	for i := 1; i < len(tab) && decroissant; i++ {
		decroissant = tab[i-1] >= tab[i]
	}
	return croissant || decroissant
}
