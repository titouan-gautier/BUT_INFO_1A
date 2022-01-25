package main

import "fmt"

func tabs(x []int, y int) int {
	return x[y]
}

func main() {
	var tint []int = []int{0, 5, 7, 8, 3, 6}
	var tab1 []int = []int{9, 8, 5, 1, 2, 6}
	var tab2 []int = []int{100, 200, 300, 400, 500}
	var n int = 3
	fmt.Println(tabs(tint, n))
	fmt.Println(tabs(tab1, n))
	fmt.Println(tabs(tab2, n))
}
