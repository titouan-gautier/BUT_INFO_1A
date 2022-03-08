class Citoyen (nom : String, prenom : String) {

    private var nom : String
    private var prenom : String
    private var numElecteur : Int
    private var bureau : BureauDeVote?

    init {
        this.nom = nom
        this.prenom = prenom
        this.numElecteur = 0
        this.bureau = null
    }

    fun modifierNumElecteur(nouveauNum : Int) {
        numElecteur = nouveauNum
    }

    fun modifierBureauDeVote(nouveauBureau : BureauDeVote) {
        bureau = nouveauBureau
    }

    fun estInscrit() = (bureau != null)

    override fun toString() : String {
        var str = "$prenom $nom"
        if (estInscrit())
            str += " (numero electeur = $numElecteur / bureau = ${bureau!!.donneNumero()})"
        return str
    }
}