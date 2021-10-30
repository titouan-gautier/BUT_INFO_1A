package nombreparfait

/*
Un nombre entier positif est dit parfait s'il est égal à la somme de
ses diviseurs propres (c'est-à-dire ses diviseurs autres que lui même)
Étant donné un nombre entier positif compris entre 0 et le plus grand
entier représentable, la fonction estParfait indique si ce nombre est
parfait ou pas.

# Entrées
- n : un nombre entier (uint64) compris entre 0 et le plus grand entier représentable par un uint64

# Sorties
- b : un booléen indiquant si n est parfait (true) ou pas (false)

# Exemples
estParfait(6) = true
estParfait(7) = false
*/

func estParfait(n uint64) (b bool) {
	var i uint64
	var nb uint64

	for i = 1; i < n; i++ {
		if n%i == 0 {
			nb = nb + i
		}
	}
	if nb > 0 {
		if n == nb {
			b = true
		} else {
			b = false
		}
	} else {
		b = false
	}
	return b
}
