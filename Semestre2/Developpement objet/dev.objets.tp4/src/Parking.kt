class Parking(raisonSociale : String, gerant : Personne, places : Int)  {
    var stationnement: Array<Voiture?>
    var raisonSociale : Int = 0
    var gerant : Personne
    var places : Int 


    init {
        this.raisonSociale = raisonSociale
        this.gerant = gerant
        this.places = places
        stationnement = arrayOfNulls<Voiture?>(places)
    }

    fun nombreDePlacesLibres() : Int {
        var count : Int = 0
        for ( i in 0 until stationnement.size ) {
            if (stationnement[i] == null){
                count++
            }
        }
        return count
    }

    fun nombreDePlacesTotales() : Int {
        return places
    }
    
    fun placeLibre(numeroPlace : Int) : Boolean {
        if (numeroPlace >= 0 && numeroPlace < stationnement.size){
            if (stationnement[numeroPlace] == null ){
                return true
            }
        }
        return false
    }

    fun stationner(numeroPlace : Int, voitureStationnee : Voiture) : Boolean {
        if (!(voitureStationnee.estGaree()) && placeLibre(numeroPlace)) {
            stationnement[numeroPlace] = voitureStationnee
            voitureStationnee.stationner(this)
            
            return true
        }
        return false
    }
}

