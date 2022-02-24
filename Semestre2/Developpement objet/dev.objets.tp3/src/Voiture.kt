class Voiture (mod : String, coul : String,vitMax : Double) {
    var modele : String = ""
    var couleur : String = ""
    var vitesseCourante : Double = 0.0
    var vitesseMaximum : Double = 0.0
    var enMarche : Boolean = false
    var proprietaire : Personne?
    var parking : Parking?

    init {
        modele = mod
        couleur = coul
        vitesseMaximum = vitMax
        proprietaire = null
        parking = null
    }

    fun estGaree() {
        parking != null
    }

    fun stationner(nouveauParking : Parking) {
        this.arreter()
        parking = nouveauParking
    }

    fun acheter(acheteur : Personne) {
        proprietaire = acheteur
    }

    fun demarrer() {
        enMarche = true
    }

    fun arreter() {
        enMarche = false
        vitesseCourante = 0.0
    }

    fun estEnMarche() : Boolean {
        if (enMarche ==  true) {
            return true
        }
        return false
    }

    fun repeindre(nouvelleCouleur : String) {
        if (enMarche == false) {
            couleur = nouvelleCouleur
        }
    }

    fun accelerer(acceleration : Double) : Double {
        if (enMarche) {
            if (acceleration >= 0.0) {
                if (vitesseCourante + acceleration <= vitesseMaximum ) {
                    vitesseCourante = vitesseCourante + acceleration
                } else {
                    vitesseCourante = vitesseMaximum
                }
            }
        }
        return vitesseCourante
    }

    fun decelerer(deceleration : Double) : Double{
        if (enMarche) {
            if (deceleration >= 0) {
                if (vitesseCourante - deceleration >= 0) {
                    vitesseCourante = vitesseCourante - deceleration
                } else {
                    vitesseCourante = 0.0
                }
            }
        }
        return vitesseCourante
    }

    fun afficher() {
        var strvoiture : String = "Voiture $modele de couleur $couleur"
        var strenmarche : String 
        if (enMarche) {
            strenmarche = "roule à $vitesseCourante / $vitesseMaximum"
        } else {
            strenmarche = "est à l'arrêt"
        }
        var strprop : String = "propriété de $proprietaire"

        if (proprietaire != null) {
            print("$strvoiture, $strprop, $strenmarche")
        } else {
            print("$strvoiture $strenmarche")
        }
    }
}