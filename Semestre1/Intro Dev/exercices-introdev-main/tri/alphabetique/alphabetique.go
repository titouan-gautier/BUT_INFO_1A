package alphabetique

/*
La fonction alphabetique trie un tableau de chaînes de caractères dans l'ordre
alphabétique.

# Entrée
- dico : le tableau de chaînes de caractères à trier
*/

func alphabetique(dico []string) {
	var i int = 1
	for i < len(dico) {
		v := dico[i]
		j := i - 1
		for j >= 0 && dico[j] > v {
			dico[j+1] = dico[j]
			j--
		}
		dico[j+1] = v
		i++
	}
}
