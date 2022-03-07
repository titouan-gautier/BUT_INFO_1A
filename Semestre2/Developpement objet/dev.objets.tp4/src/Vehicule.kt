class Vehicule(mod : String, coul : String,vitMax : Double)  {
    var modele : String = ""
    var couleur : String = ""
    var vitesseCourante : Double = 0.0
    var vitesseMaximum : Double = 0.0
    var enMarche : Boolean = false
    
    init {
        modele = mod
        couleur = coul
        vitesseMaximum = vitMax
    }

    fun demarrer() {
        enMarche = true
    }

    fun arreter() {
        enMarche = false
        vitesseCourante = 0.0
    }

    fun estEnMarche() : Boolean {
        return enMarche
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
}