abstract class Animal(nom : String, age : Int) : Appelable {
  protected var nom :String 
  protected var age : Int 
  protected var maitre : Personne? = null

  init {
      this.nom = nom
      this.age = age
  }

  override fun repondre(unNom : String) =
          (nom == unNom)
 
  fun estAdoptePar(p : Personne) {
    maitre = p
  }

  abstract fun ageHumain() : Int

  open fun courir() {
    println("$nom court !!!!")
  }
}
