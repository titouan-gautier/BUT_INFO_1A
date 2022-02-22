class BureauDeVote (num : Int, maxElecteurs : Int) {

    private val numero : Int
    private var president : Citoyen?
    private val listeElectorale : Array<Citoyen?>

    init {
      numero  = num
      president = null
      listeElectorale = arrayOfNulls<Citoyen>(maxElecteurs)
    }

    fun designerPresident(pers : Citoyen) =
      if (pers.estInscrit()) {
        president = pers
        true
      }
      else false

    fun electionPossible() =
      (president != null 
       && listeElectorale.isNotEmpty())

    private fun dernierePosition() : Int {
        for (i in 0 until listeElectorale.size) {
            if (listeElectorale[i] == null)
            return i
        }
        return listeElectorale.size
    }

    fun inscrire(nouveau : Citoyen) : Boolean {
        if (nouveau.estInscrit())
            return false // nouveau deja inscrit
        val dernier = dernierePosition()
        if (listeElectorale.size == dernier)
            return false // listeElectorale pleine
        listeElectorale.set(dernier, nouveau)
        nouveau.modifierBureauDeVote(this)
        nouveau.modifierNumElecteur(dernier+1)
        return true
    }

    fun reorganiserListe() {
        // TODO
    }
    /*
        listeElectorale.sort() // KO : on y reviendra
        for (i in 0 until dernierePosition()) {
            listeElectorale[i]?.modifierNumElecteur(i+1)
        }
    }
    */

    override fun toString() : String {
        var str = "Bureau $numero"
        if (president != null)
            str += " présidé par $president"
        str += " liste = ["
        for (i in 0 until dernierePosition()) {
            str+= " ${listeElectorale[i]?.toString()}"
        }
        str += " ]"
        return str
    }

    fun donneNumero() = numero
}