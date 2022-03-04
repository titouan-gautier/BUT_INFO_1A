class Personne(nom : String, prenom : String) {
    var nom : String = ""
    var prenom : String = ""

    init {
        this.nom = nom
        this.prenom = prenom
    }

    fun donneNomComplet() : String {
        var monNom = nom.uppercase()
        return "$prenom $monNom"
    }
}