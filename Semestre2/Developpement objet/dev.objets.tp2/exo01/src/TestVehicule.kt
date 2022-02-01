
fun main() {
    var maSaxo = Vehicule("Saxo","rouge",4,180.0)
    println(maSaxo)
    maSaxo.demarrer()
    println(maSaxo)
    maSaxo.accelerer(50.0)
    println(maSaxo)
    maSaxo.accelerer(40.0)
    println(maSaxo)
    println(maSaxo.vitesse())

    var maYaris = Vehicule("Yaris","grise",4,150.0)
    maYaris.demarrer()
    maYaris.accelerer(150.0)
    println(maYaris)

    println(maYaris.vaPlusVite(maSaxo))

    maSaxo.arreter()
    maYaris.arreter()
}
