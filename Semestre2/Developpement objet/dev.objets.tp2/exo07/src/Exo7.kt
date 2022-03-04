/*
La fonction somme doit calculer la somme des nombres contenus dans un tableau
d'entiers puis retourner cette somme. Il faudra bien penser à se demander ce
qu'est la somme d'un tableau vide.
@author init.dev (L.Jezequel)
@param tab : un tableau d'entiers complètement remplis
@return un entier correspondant à la somme des éléments de tab
*/

fun somme(tab : Array<Int>) : Int {
    var c : Int = 0
    for (i in 0 until tab.size) {
        c += tab[i]
    }
    if (c == null) {
        return -1
    }
    return c
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
Étant donné un tableau, la fonction envers doit l'inverser (en place), c'est à
dire que le premier élément du tableau doit devenir le dernier, le deuxième
élément doit devenir l'avant dernier, etc jusqu'au dernier élément qui doit
devenir le premier.
La fonction modifie le tableau d'entrée et n'a donc pas de sorties.
@param tab : le tableau à inverser
*/

fun envers(tab : Array<Int>) {
    for (i in 0 until tab.size/2) {
        var a : Int = tab[i]
        var b : Int = tab[tab.size-i - 1]
        tab[i] = b
        tab[tab.size - i - 1] = a
    }
}

/*
La fonction alphabetique trie un tableau de chaînes de caractères dans l'ordre
alphabétique.
@author init.dev (L.Jezequel)
@param dico : le tableau de chaînes de caractères à trier
*/

fun alphabetique(dico : Array<String>) {
    dico.sort()
}

/*
La fonction décroissant doit trier un tableau d'entiers du plus grand
au plus petit.
@author init.dev (L.Jezequel) 
@param tab : le tableau à trier
*/
fun decroissant(tab : Array<Int>) {
    tab.sortDescending()
}

/* 
La fonction donne le plus grand élément contenu dans le tableau. Null si le tableau est vide
@param tab : le tableau dans lequel chercher
@return  la plsu grande valeur contenue dans tab (si elle existe), null sinon
*/
fun lePlusGrand(tab : Array<Int>) : Int? {
    var max : Int = 0
    if (tab.size == 0) {
        return null
    }
    for (i in 0 until tab.size) {
        if (tab[i] > max) {
            max = tab[i]
        }
    }
    return max
}

/* 
La fonction donne la valeur la plus petite, mais plus grande que v. 
Null si le tableau est vide ou si v n'a pas de plus grande valeur dans tab
v n'est pas obligatoirement dans tab, ca ne change rien
@param tab : le tableau dans lequel chercher
@param v : la valeur à rechercher
@return  la valeur juste plus grande que v dans tab (si elle existe), null sinon
*/

fun justePlusGrandeQue(tab : Array<Int>, v : Int) : Int? {
    var max : Int = lePlusGrand(tab)!!
    var index : Int = recherche(tab, v)
    if (v == lePlusGrand(tab)) {
        return null
    }
    if (index > 0) {
        if (tab[index] > max ) {
            return null
        }
    }
    for (i in 0 until tab.size) {
        if (tab[i] > v) {
            if (tab[i] < max) {
                max = tab[i]
            }
        }
    }
    return max

}


