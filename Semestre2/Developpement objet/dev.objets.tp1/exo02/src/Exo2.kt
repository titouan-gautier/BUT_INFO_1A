
const val MAX = 10_000
var x : Double = 0.0
var y = 12
var z : Boolean = true

fun main() {
    x = 2.0
    y++
    truc(y)
    z = (x == 2.0)
    if (inc(x) == 3.0)
        print("ok")
}

fun truc(a : Int) {
    println(a)
}

fun inc(a : Double) = a + 1

