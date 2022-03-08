class Chat(nom : String, age : Int, pedigree : String) 
      : Animal(nom, age) {
  
  private val pedigree : String 

  init {
      this.pedigree = pedigree
  }

  fun miauler() {
    println("$nom dit : miaou, miou...") 
  }

  override fun ageHumain() : Int {
    return age * 6
    // https://www.santevet.com/articles/comment-estimer-l-age-de-mon-chat-en-age-humain
  }
}