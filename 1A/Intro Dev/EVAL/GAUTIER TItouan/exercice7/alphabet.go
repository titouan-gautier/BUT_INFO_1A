package alphabet

/*
La fonction alphabet doit indiquer si deux mots sont dans l'ordre alphabétique.
On considérera que les mots en questions ne sont constitués que de lettres
minuscules non accentuées, donc prises dans l'ensemble {a, b, c, d, e, f, g, h,
i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z}.

# Entrées
- m1 : un mot constitué de lettres minuscules non accentuées
- m2 : un mot constitué de lettres minuscules non accentuées

# Sorties
- m1avantm2 : un booléen qui vaut true si m1 est (strictement) avant m2 dans
l'ordre alphabétique et qui vaut false sinon

# Indication
On peut comparer les caractères d'une chaîne de caractères comme on compare des
entiers. Ainsi, m1[i] < m2[j] vaut true si et seulement si le i-ième caractère
de m1 est avant le j-ième caractère de m2 dans l'alphabet.
*/

func alphabet(m1, m2 string) (m1avantm2 bool) {
	var tab string = ""
	if m1 == m2 {
		return false 
	}
	if m1 != tab {
		if m2 != tab {
			for i := 0 ; i<len(m1) ; i++{
				if m1[i] < m2[i]  {
					return true
				}
			}
		} else {
			return false
		}
	} else {
		return true
	}
	return true
}