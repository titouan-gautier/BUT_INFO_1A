fun main() {
    var ab : String
    var yz : String 
    ab = "abcd"
    yz = "ABCD"
    println("val: ${ab == yz}") // false
    println("ref: ${ab === yz}")// false
    yz = ab                 // ij="abcd"
    println("val: ${ab == yz}") // true
    println("ref: $``{ab === yz}")// true
    var ij = "ABCD"
    yz = ij.lowercase() // ij="abcd"
    println("val: ${ab == yz}") // true
    println("ref: ${ab === yz}")// false
}