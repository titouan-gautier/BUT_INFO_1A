/*
La fonction perimetre doit retourner le perimetre d'un rectangle
 defini par sa longueur et sa largeur.
 
 @param longueur la longueur du rectangle
 @param largeur la largeur du rectangle
 
 @return le perimètre du rectangle
  */
fun perimetre(longueur : Double, largeur : Double) : Double {
    return 2*longueur + 2*largeur
}

/*
La fonction sommeNombresPairs doit retourner la somme de tous les nombres
pairs contenus entre deux entiers (inclus).

@author init.dev (L.Jezequel)

@param min  un nombre entier, une des bornes
@param max  un nombre entier, l'autre borne

@return la somme de tous les nombres pairs compris entre min et max inclus

# Exemple
sommeNombresPairs(2, 7) = 2 + 4 + 6 = 12
*/
fun sommeNombresPairs(min : Int, max : Int) : Int {
    var somme : Int = 0
    for (i in min until max+1) {
        if (i%2 == 0) {
            somme +=i
        }
    }
    return somme
}

/*
La fonction compte doit indiquer combien de fois une valeur v apparaît dans un
tableau tab.

@author init.dev (L.Jezequel)

@param tab un tableau d'entiers
@param v la valeur à chercher

@return le nombre de fois que v apparaît dans tab
*/

fun compte(tab : Array<Int>, v : Int) : Int {
    var count : Int = 0
    for (i in 0 until tab.size) {
        if (tab[i] == v) {
            count += 1
        }
    }
    return count
}

/*
La fonction bienrange indique si un tableau d'entiers est bien trié par ordre
croissant ou pas.

@author init.dev (L.Jezequel)

@param tab le tableau d'entiers à analyser

@return un booléen qui vaut true si les entiers de tab sont bien triés en
ordre croissant et false s'ils ne sont pas bien triés.
*/

fun bienrange(tab : Array<Int>) : Boolean? {
    for (i in 1 until tab.size) {
        if (tab[i-1] > tab[i]) {
            return false
        }
    }
    return true
}

/* 
La fonction filtreEtmultiplie parcours le tableau et multiplie les éléments non null.

@param tab un tableau de doubles possiblement null 
     NB : tab contiendra toujours au moins un élément non null

@return le produit des éléments non-null du tableau
*/

fun filtreEtMultiplie(tab : Array<Double?>) : Double {
    var produit : Double = 0.0
    var c : Int = 0
    for (i in 0 until tab.size) {
        if (tab[i] != null) {
            c++
            if (c == 1) {
                produit = tab[i]!!.toDouble()
            }
            else {
                produit = produit*tab[i]!!.toDouble()
            }
            
        }
    }
    return produit
}