package chocolats

import "fmt"

/*
Une marque de barres de chocolat fait une promotion~: si on retourne k embalages on en obtient une gratuite. On se demande alors combien de barres de chocolat on peut obtenir quand on dispose de n euros et que chacune coûte m euros.
La fonction miam doit répondre (de manière récursive) à cette question.

# Entrées
- n : la somme dont on dispose en euros
- m : le prix d'une barre de chocolat en euros, toujours supérieur à 0
- k : le nombre d'embalages qu'il faut pour obtenir une barre de chocolat gratuite, toujours supérieur à 0

# Sortie
- choco : le nombre de barre de chocolat qu'on peut obtenir en tout

# Exemple

k = nombre d'emballages pour un gratuit
n = argent que l'on possède
m = prix d'une barre
e = nombre d'mballage qu'on a
choco = nombre de barre achteté
*/

func miam(n, m, k uint) (choco uint) {
	var e uint
	if n < m || m == 0 {
		return 0
	}
	for i := n; i <= 0; i = i - m {
		n = n - m
		choco = choco + 1
		fmt.Println(n)
		e = e + 1
		if e == k {
			k = k + 1
			e = 0
		}
		if i <= 0 {
			choco = choco + k
		} else {
			miam(n, m, k)
		}
	}
	return choco
}
