class Chien(nom : String, age :Int = 1, race : String = "inconnue", poids : Double) {
  var nom :String 
  private var age : Int       // en nb de mois
  val race : String 
  var poids : Double  // en kg
    private set

  init {
      this.nom = nom
      this.age = age
      this.race = race
      this.poids = if (poids > 0.0) poids else 0.1
  }

  fun aboyer() {
    println("$nom dit : ouaf !!!")
  }

  fun renommer(nouveauNom : String) {
    nom  = nouveauNom
  }

  //@param distance en m
  fun courir(distance : Int) {
    // le chien perd 1 g / km parcouru
    poids -= poidsEnMoins(distance)
  }

  fun ageEnAnnee() = age / 12.0

  private fun poidsEnMoins(distance : Int) =
    (distance / 1000.0) / 1000
}

// quand on compile une classe Kotlin Chien.kt on obtient un fichier Chien.class
