
/*
La fonction somme doit calculer la somme des nombres contenus dans un tableau
d'entiers puis retourner cette somme. Il faudra bien penser à se demander ce
qu'est la somme d'un tableau vide.

@author init.dev (L.Jezequel)

@param tab : un tableau d'entiers complètement remplis

@return un entier correspondant à la somme des éléments de tab

*/

fun somme(tab : Array<Int>) : Int {
    var sum : Int = 0
    for (i in 0 until tab.size) {
        sum = sum + tab[i]
    } 
    return sum
}

/*
La fonction recherche doit indiquer la premiere position d'une valeur v dans un tableau
tab. Si la valeur v n'est pas présente, elle indique -1

@author init.dev (L.Jezequel)

@param tab : le tableau dans lequel chercher
@param v : la valeur à chercher

@return  la position de v dans tab (si elle existe)

*/

fun recherche(tab : Array<Int>, v :Int) : Int {
    for (i in 0 until tab.size) {
        if (tab[i] == v) {
            return i
        }
    }
    return -1
}

/*
Un ensemble d'entier est un paquet de plusieurs entiers, sans doublons.
La fonction estEnsemble doit indiquer si en tableau d'entiers correspond à un
ensemble ou non.

@author init.dev (L.Jezequel)

@param E : un tableau d'entiers

@return un booléen indiquant si E est bien un ensemble ou non 

*/
fun estEnsemble(tab : Array<Int>) : Boolean {
    for (i in 0 until tab.size) {
        for (j in i+1 until tab.size) {
            if (tab[j] == tab[i]) {
                return false
            }
        }
    }
    return true
}

/*
Étant donné un tableau, la fonction envers doit l'inverser (en place), c'est à
dire que le premier élément du tableau doit devenir le dernier, le deuxième
élément doit devenir l'avant dernier, etc jusqu'au dernier élément qui doit
devenir le premier.

La fonction modifie le tableau d'entrée et n'a donc pas de sorties.

@author init.dev (L.Jezequel)

@param tab : le tableau à inverser

*/

fun envers(tab : Array<Int>) {
    for (i in 0 until tab.size/2) {
        var a : Int =tab[i]
        var b : Int = tab[(tab.size-1)-i]
        tab[i] = b
        tab[(tab.size-1)-i] = a
    }
}


