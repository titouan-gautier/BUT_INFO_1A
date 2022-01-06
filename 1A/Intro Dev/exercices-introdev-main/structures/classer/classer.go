package classer

/*
La fonction classer doit trier un tableau d'étudiants (tels que définis) ci-dessous
de celui qui a la meilleur moyenne (la plus élevée) à celui qui a la plus mauvaise
moyenne. En cas d'égalité de moyenne entre deux étudiants, celui qui a le nom de
famille qui arrive en premier dans l'ordre alphabétique doit être classer avant
l'autre (on peut utiliser <, >, <=, >=, == pour comparer les chaînes de caractères
par ordre alphabétique). Si deux étudiants ont la même moyenne et le même nom, on
met en premier celui qui a le prénom qui est en premier dans l'ordre alphabétique.

# Entrée
- promo : le tableau d'étudiants à trier

# Sortie
- classement : un tableau contenant les mêmes étudiants mais trié
*/

type etudiant struct {
	nom, prenom string
	moyenne     float64
}

func classer(promo []etudiant) (classement []etudiant) {
	if promo == nil && len(promo) == 0 {
		return nil
	}
	for i := 0 ; i < len(promo) ; i++ {
		classement = append(classement, promo[i])
		var j etudiant = 0
		var v etudiant = promo[i]
		for j < len(classement) && classement[j] <= v {
			j++
		}
		for j < len(trier) {
			tmp := trier[j]
			trier[j] = v
			v = tmp
			j++
		}
	}
	return classement
}
