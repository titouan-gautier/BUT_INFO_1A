class VoitureTurbo(Voiture : Voiture, mod : String, coul : String, vitMax: Double)  {
    var Voiture : Voiture
    private var turbo : Boolean = false

    init {
        modele = mod
        couleur = coul
        vitesseMaximum = vitMax
        this.Voiture = Voiture
    }

    fun changeTurbo(etat : Boolean) {
        this.turbo = true
    }

    fun accelerer(acceleration : Double) : Double {
        if (this.turbo == true) {
            this.Voiture.acceleration *= 3
        }
        return Voiture.acceleration
    }
}