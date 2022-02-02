fun main() {
    var ab : String? = "abc"
    var yz : String? = "ABC"
    println(ab)
    println(yz)
    println("-----")
    yz = ab
    println(ab)
    println(yz)
    println("-----")
    ab = null
    println(ab)
    println(yz)
    println("-----")
}