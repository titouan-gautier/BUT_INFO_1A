fun main() {
    val anti = "anticonstitutionnellement"
    println(anti.length)
    println(anti[2])
    println(anti[4])
    println(anti[19])
    println(anti.reversed())
    println(anti.substring(1,4))
    println(anti.lastIndexOf("n"))
    println(anti.indexOf("n"))

    var bonjour = "bonjour"
    println(bonjour)
    val btlm = bonjour + " tout" + " le monde"
    println(btlm)
    var b : String = "bonjour"

    for (i in 0 until 25) {
        b = b + b.random()
    }
    println(b)
    var majb = bonjour.uppercase()
    var minb = majb.lowercase()
    println("$majb , $minb")

    println(majb.equals(minb))
    println(majb.compareTo(minb))
    println(majb == minb)
    println(majb === minb)
    println(majb <= minb)
    println(majb >= minb)
    

}
