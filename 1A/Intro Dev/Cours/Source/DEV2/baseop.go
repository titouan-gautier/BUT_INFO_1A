package baseop

func Plus(x int, y int) int {
	var res int = x
	for i := 0; i < y; i++ {
		res++
	}
	return res
}

func Mult(x int, y int) int {
	return 0
}
