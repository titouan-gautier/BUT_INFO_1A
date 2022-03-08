fun main() {
    var rogue = Chien("Rogue", 15, "Berger Australien")
    var potter = Chien("Potters", 40, "Beauceron")
    var gaga = Chat("Gaga", 88, "de maison")
    rogue.aboyer()
    potter.aboyer()
    // potter.miauler() : unresolved reference: miauler
    gaga.miauler()
    rogue.repondre("Potter")
    rogue.repondre("Rogue")

    var animal : Animal = Chien("Rogue", 16, "Berger Australien")
    animal.repondre("Potter")
    //animal.aboyer()
    animal = Chat("Gaga", 89, "de maison")
    animal.repondre("Potter")

    val arnaud = Personne("Arnaud", "Lanoix Brauer")
    arnaud.adopte(rogue)
    rogue.estAdoptePar(arnaud)
    arnaud.adopte(gaga)
    arnaud.repondre("Arnaud")
    rogue.rapporte("balle")
}