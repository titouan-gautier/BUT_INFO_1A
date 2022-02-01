/*
On considère des ensembles de nombres représentés par des tableaux : on considère
que ces tableaux ne contiennent qu'une seule fois chaque nombre (puisqu'ils
représentent des ensembles) et les nombres ne sont pas nécessairement dans
l'ordre dans les tableaux.

On veut savoir si deux ensembles sont égaux ou pas, c'est-a-dire savoir si deux
tableaux contiennent les mêmes nombres ou pas. C'est à cette question que doit
répondre la fonction egalite.

@author init.dev (L.Jezequel)

@param tab1 un tableau d'entiers (sans doublons) représentant un ensemble
@param tab2  un tableau d'entiers (sans doublons) représentant un ensemble

@return un booléen qui vaut true si tab1 et tab2 représentent le même ensemble et
          qui vaut false sinon
*/

fun egalite(tab1 : Array<Int>, tab2 : Array<Int>) : Boolean {
    var count : Int = 0
    if (tab1.size != tab2.size) {
        return false
    }
    for (i in 0 until tab1.size) {
        for (j in 0 until tab2.size) {
            if (tab1[i] == tab2[j]) {
                count = count + 1
            }
        }
    }
    if (count == tab1.size) {
        return true
    }
    return false;
}