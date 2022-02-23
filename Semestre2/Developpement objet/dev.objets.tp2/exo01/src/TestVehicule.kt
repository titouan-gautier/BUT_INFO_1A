
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

    var voiture1 = Vehicule("V1","couleur1",4,130.0)
    var voiture2 = Vehicule("V2","couleur2",4,130.0)
    var voiture3 = Vehicule("V2","couleur3",4,130.0)
    var voiture4 = Vehicule("V3","couleur4",4,130.0)

    var camion = arrayOf(voiture1,voiture2,voiture3,voiture4)
    print(afficherC(camion))
}

fun afficherC(camion : Array<Vehicule>) {
    for (i in camion.indices) {
        println(camion[i])
    }
}
