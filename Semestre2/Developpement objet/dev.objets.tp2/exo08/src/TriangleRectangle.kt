fun main(args : Array<String>) {
	// TODO
}

fun estRectangle(ab : Double, ca : Double, cb : Double) : Boolean {
	var abcarre : Double = ab^2.toInt()
	var cacarre : Double = ca^2.toInt()
	var cbcarre : Double = cb^2.toInt()

	if (abcarre == cacarre + cbcarre) {
		return true
	}
	return false
}