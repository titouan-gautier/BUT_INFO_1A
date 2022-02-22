fun main() {
    var rogue = Chien("Rogue", 15, "Berger Australien", 30.2)
    var potter = Chien("Other", 40, "Beauceron", 39.4)
    var soukie = Chien("Soukie", poids = 4.7)

    rogue.aboyer()
    potter.renommer("Potter")
    println("${potter.nom} : ${potter.poids} kg")
    potter.courir(400_000)
    println("${potter.nom} : ${potter.poids} kg")
    var age = rogue.ageEnAnnee()
    println("${rogue.nom} : $age ans")
    soukie.aboyer()

    println("${rogue.nom}")
    rogue.nom = "Severus"
    //println("${rogue.age}")   // error
      // it is private in 'Chien'
    //rogue.age = 10            // error 
      // it is private in 'Chien' 
    println("${rogue.race}")
    //rogue.race = "Serpentard" // error
      // val cannot be reassigned
    println("${rogue.poids}")
    //rogue.poids = 42.0        // error
      // the setter is private in 'Chien'
    rogue.courir(100)
    //rogue.poidsEnMoins(100) // error
      // it is private in 'Chien'

}