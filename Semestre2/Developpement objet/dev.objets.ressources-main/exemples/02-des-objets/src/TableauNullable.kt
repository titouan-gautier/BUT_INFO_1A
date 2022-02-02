fun main() {
    var ta : Array<Int?>
    ta = arrayOf(3, 5, 7)
    var tz = ta  
    afficher(ta)
    afficher(tz)
    ta[2] = 2
    tz[1] = 4
    tz[0] = null
    afficher(ta)
    afficher(tz)
    ta = arrayOfNulls<Int>(4)
    afficher(ta)
    afficher(tz)
    ta[3] = tz[2]
    ta[2] = tz[1]
    afficher(ta)
    afficher(tz)
}

fun afficher(t : Array<Int?>) {
    print("[ ")
    for (i in t) {
        print("$i ")
    }
    println("]")
}