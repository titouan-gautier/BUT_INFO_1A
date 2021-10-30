package main

import "fmt"

func calcul (x int, y int, t [][]int) [][]int {
	for i:= 0; i < x; i++ {
		for j := 0; j < y ; j++ {
			t[i][j] = (i+1)*j
		}
	}
	return t
}
func main() {
	var m int = 10
	var n int = 11
	var tab [][]int = make([][]int, m)
	for i:= 0; i < m; i++ {
		tab[i] = make([]int, n)
	}
	fmt.Println(calcul(m, n, tab))
}
