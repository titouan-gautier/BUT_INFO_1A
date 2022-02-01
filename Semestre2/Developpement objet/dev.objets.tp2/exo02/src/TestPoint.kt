
fun main() {
	var p1 = Point()
	var p2 = Point(1,2)
	println("$p1,$p2")
	println(p1.getX())
	p2.setY(3)
	println(p2.getY())
	p2.translater(1,1)
	println(p2)
	p1.deplacer(2,10)
	println(p1)
	println(p2.distanceOrigine())
	println(p1.distance(p2))
	p1 = p2
	p1.translater(2,1)
	println(p1)
	println(p2)
}
