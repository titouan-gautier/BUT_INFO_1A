class Chien(nom : String, age : Int, race : String) 
      : Animal(nom, age),  Joueur {
  private val race : String 

  init {
      this.race = race
  }

  fun aboyer() {
    println("$nom dit : ouaf ouaf !!!") 
  }

  override fun ageHumain() : Int {
    return age * 7
    // https://www.santevet.com/articles/comment-estimer-l-age-de-mon-chien-en-age-humain
  }

  override fun courir() {
    aboyer()
    super.courir()
    aboyer()
    aboyer()
  }

  override fun rapporte(objet : String) {
    courir()
    print("$nom rapporte $objet")
    if (maitre != null)
      print(" à ${maitre!!.donneNomComplet()}")
    println("")
  }
}