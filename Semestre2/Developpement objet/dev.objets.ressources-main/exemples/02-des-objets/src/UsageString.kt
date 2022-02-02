fun main() {

    var chaine : String = "ToToRo"
    println(chaine)
    println("----\nproprietes\n ----")
    val longueur = chaine.length
    println("longueur = $longueur") 
    var indice = chaine.lastIndex
    println("dernier indice = $indice")

    println("----\nvide/pas vide\n ----")
    var verif : Boolean
    verif = chaine.isEmpty() 
    print(verif)
    verif = chaine.isNotEmpty()
    print(verif) 
    verif = chaine.isNotBlank()
    println(verif) 

    println("----\nacces indices\n ----")
    var carac = chaine.get(0)
    println(carac)
    carac = chaine[0]
    println(carac)
    carac = chaine[3]
    println(carac)

    println("----\nconcatenations\n ----")
    var chaine2 : String
    chaine2 = "Tonari no ".plus(chaine)
    println(chaine2)
    chaine2 = "Tonari no " + chaine
    println(chaine2)

    println("----\nappartients\n ----")
    verif = chaine.contains("o")
    print(verif)
    verif = "o" in chaine
    print(verif)
    verif = chaine.contains("O")
    print(verif)
    verif = chaine.contains("O", 
                ignoreCase = true)
    println(verif)

    println("----\negalité\n ----")
    verif = chaine.equals("ToToRo")
    print(verif)
    verif = chaine == "ToToRo"
    print(verif)
    verif = chaine == "totoro"
    print(verif)
    verif = chaine != "totoro"
    print(verif)
    verif = chaine.equals("totoro", 
                ignoreCase = true)
    println(verif)


    println("----\ncomparaison\n ----")
    var comp = chaine.compareTo("TOTORO")
    println(comp)
    comp = chaine.compareTo("totoro")
    println(comp)
    verif = chaine > "totoro"
    println(verif)
    verif = chaine <= "totoro"
    println(verif)

    println("----\nminuscule/majuscule\n ----")
    chaine2 = chaine.uppercase()
    println(chaine2)
    chaine2 = chaine.lowercase()
    println(chaine2)


    println("----\nrenverse\n ----")
    chaine2 = chaine.reversed()
    println(chaine2)

    println("----\nrandom\n ----")
    carac = chaine.random()
    println("carac aléa : $carac")
    carac = chaine.random()
    println("carac aléa : $carac")
    carac = chaine.random()
    println("carac aléa : $carac")

    // here
    println("----\nsous-chaines\n ----")
    chaine2 = chaine.substring(2)
    println(chaine2)
    chaine2 = chaine.substring(2, 4)
    println(chaine2)

    println("----\nindices de\n ----")
    indice = chaine.indexOf('o')
    println("indice de o = $indice")
    indice = chaine.lastIndexOf('o')
    println("dernier indice de o = $indice")
    indice = chaine.indexOf('O')
    println("indice de O = $indice")
    indice = chaine.indexOf('O', ignoreCase=true)
    println("indice de O = $indice")






}