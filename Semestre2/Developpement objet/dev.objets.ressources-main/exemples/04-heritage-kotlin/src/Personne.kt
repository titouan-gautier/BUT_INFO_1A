class Personne (prenom:String, nom:String) : Appelable {
  private val nom : String
  private val prenom : String
  private val compagnons : Array<Animal?>
  private var nbCompagons : Int = 0

  init {
      this.nom = nom
      this.prenom = prenom
      this.compagnons = arrayOfNulls<Animal>(10)
  }

  fun adopte(compagnon : Animal) {
    if (nbCompagons < compagnons.size) {
        compagnons[nbCompagons] = compagnon
        nbCompagons++
    }
  }

  override fun repondre(unNom : String) =
          (nom + " " + prenom == unNom)

  fun donneNomComplet() = "$prenom ${nom.uppercase()}"
}