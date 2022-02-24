

import kotlin.arrayOfNulls

class Parking(places : Int) {
    var stationnement: Array<Voiture?>
    var places : Int

    init {
        this.places = places
        stationnement = arrayOfNulls<Voiture?>(places)
    }

    fun nombreDePlacesLibres() : Int {
        var c : Int = 0
        for (i in 0 until stationnement.size ) {
            if (stationnement[i] != null) {
                c += 1
            }
        }
        return c
    }

    fun nombreDePlacesTotales() : Int {
        return places
    }
    
    fun placeLibre(numeroPlace : Int) : Boolean {
        if (stationnement[numeroPlace] != null) {
            return false
        }
        return true
    }

    fun stationner(numeroPlace : Int, voitureStationnee : Voiture) : Boolean {
        if (!placeLibre(numeroPlace)) {return false}
        
        stationnement[numeroPlace] = voitureStationnee
        voitureStationnee.stationner(this)
        return true
    }
}