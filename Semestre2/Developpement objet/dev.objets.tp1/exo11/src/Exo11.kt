/*
La fonction décroissant doit trier un tableau d'entiers du plus grand
au plus petit.

@author init.dev (L.Jezequel) 

@param tab : le tableau à trier
*/

fun decroissant(tab : Array<Int>) {
    for (i in 0 until tab.size) {
        var x : Int = tab[i]
        var j : Int = i
        while (j > 0 && tab[j-1] < x) {
            tab[j] = tab[j-1]
            j = j - 1
        }
        tab[j] = x
    }
    // https://fr.wikipedia.org/wiki/Tri_par_insertion
}

/*
La fonction alphabetique trie un tableau de chaînes de caractères dans l'ordre
alphabétique.

@author init.dev (L.Jezequel)

@param dico : le tableau de chaînes de caractères à trier
*/

fun alphabetique(dico : Array<String>) {
    for ( i in dico.size-1 downTo 1 ) {
        for (j in 0 until i) {
            if (dico[j+1] < dico[j]) {
                var a = dico[j+1]
                var b = dico[j]
                dico[j+1] = b
                dico[j] = a
            }
        }
    }
    //https://fr.wikipedia.org/wiki/Tri_%C3%A0_bulles
    // NB : Les opérateurs de comparaison fonctionnent sur les chaines de caractères
}

