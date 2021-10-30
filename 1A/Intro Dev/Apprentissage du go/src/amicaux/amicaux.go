package main

import "fmt"

func div1(x int, s1 int, s2 int) int {
	for i := 1; i < x; i++ {
		if x%i == 0 {
			s1 = i + s1
		}
	}
	return s1
}

func div2(y int, s1 int, s2 int) int {
	for i := 1; i < y; i++ {
		if y%i == 0 {
			s2 = i + s2
		}
	}
	return s2
}

func amicaux(x int, y int, n int, m int) bool {
	return x == m && y == n
}

func main() {
	var n int = 220
	var m int = 284
	var somme1 int
	var somme2 int
	var a1 int = (div1(n, somme1, somme2))
	var a2 int = (div2(m, somme1, somme2))
	fmt.Println(amicaux(a1, a2, n, m))
}
