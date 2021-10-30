package maxoccurrences

/*
Étant donné un tableau d'entiers t, la fonction maxoccurrences doit retourner
l'entier qui apparaît le plus souvent dans t et le nombre de fois que cet
entier apparaît. En cas d'égalité on choisira arbitrairement l'un des entiers
à égalité. Si le tableau est vide on retournera un entier quelconque et 0 pour son nombre d'apparitions.

# Entrées
- t : le tableau dans lequel chercher

# Sortie
- n : l'entier qui apparaît le plus de fois dans t
- occ : le nombre de fois que n apparaît dans t

# Exemple
maxoccurrences([]int{1, 2, 3, 4, 3}) = 3, 2
*/
func maxoccurrences(t []int) (n int, occ int) {
	var occmax int = 0
	for i := 0; i < len(t); i++ {
		occ = 0
		for k := 0; k < len(t); k++ {
			if t[i] == t[k] {
				occ = occ + 1
			}
		}
		if occmax < occ {
			occmax = occ
			n = t[i]
		}
	}
	return n, occmax
}
