class Camion(places : Int) {
    private var placesOccupees : Int = 0
    private var places : Int
    var remorque : Array<Voiture?>

    init {
        this.places = places
        remorque = arrayOfNulls<Voiture?>(places)
    }

    fun estVide() : Boolean {
        if (placesOccupees == 0) {
            return true
        }
        return false
    }
    fun estPlein() : Boolean{
        if (placesOccupees == places) {
            return true
        }
        return false
    }

    fun charger(voitureTransportee : Voiture) : Boolean {
        if (!this.estPlein() && !voitureTransportee.estEnMarche() && voitureTransportee.camion == null) {
            remorque[placesOccupees] = voitureTransportee

            placesOccupees++

            voitureTransportee.charger(this)

            return true
        }
        return false
    }

    fun decharger() : Voiture? {
        if (!estVide()){

            var voitureCharger : Voiture? = remorque[placesOccupees-1]

            remorque[placesOccupees-1]?.decharger()

            remorque[placesOccupees-1] = null

            placesOccupees--

            return voitureCharger

        }

        return null

    }

}

