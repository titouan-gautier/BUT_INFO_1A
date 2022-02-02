fun main() {
    var w : Int
    var x : Int?
    var y : Double? = 10.0
    var z : String? = "totoro"
    //w = null 
    // erreur de compilation
    y = null
    //z = null
    //val l = z.length 
    // erreur de compilation

    var l = z?.length
    println(l)

    l = z!!.length
    println(l)

    l = z?.length ?: 0
    println(l)
}