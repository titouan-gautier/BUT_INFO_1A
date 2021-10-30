package recherche

/*
La fonction recherche doit indiquer la position d'une valeur v dans un tableau
tab. Si la valeur v n'est pas présente, il faut l'indiquer en retournant false.

# Entrées
- tab : le tableau dans lequel chercher
- v : la valeur à chercher

# Sorties
- found : un booléen indiquant si v est présente dans tab ou pas
- pos : la position de v dans tab (si elle existe)

*/

func recherche(tab []int, v int) (trouve bool, pos int) {
	if tab != nil {
		for i := 0 ; i<len(tab); i++ {
			if tab[i] == v {
				return true , i
			}
		}
	} else {
		return false, 0
	}
	return trouve, pos
}
