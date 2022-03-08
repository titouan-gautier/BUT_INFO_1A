import kotlin.random.*

fun main() {
	var p1 = Point()
	var p2 = Point(1,2)
	/* println("$p1,$p2")
	println(p1.getX()) */
	p2.setY(3)
	/* println(p2.getY()) */
	p2.translater(1,1)
	/* println(p2) */
	p1.deplacer(2,10)
	/* println(p1)
	println(p2.distanceOrigine())
	println(p1.distance(p2)) */
	p1 = p2
	p1.translater(2,1)
	/* println(p1)
	println(p2)
 */
	var tab : Array<Point?> = arrayOfNulls(10)
	tab[0] = p1

	/* print(afficher(tab)) */

	remplir(tab)

	afficher(tab)

}

fun afficher(points : Array<Point?> ) {
	for (i in 0 until points.size) {
		if (points[i] != null) {
			println(points[i])
		}
	}
}

fun remplir(points : Array<Point?>) {
	for (i in 0 until points.size) {
		var dx = Random.nextInt(1,10)
		var dy = Random.nextInt(1,10)
		if (points[i] == null) {
			if (i == 0) {
				points[i] = Point(dx,dy)
			}
			else {
				var p = points[i-1]
				points[i] = p?.translater(dx,dy)
				
			}
		}
	}
}