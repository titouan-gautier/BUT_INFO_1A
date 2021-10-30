package main

import "fmt"

func tab (x int) []int {
	var res []int = make([]int, x)
	for i:= 0; i < x; i++ {
		res[i] = i*7
	}
	return res
}

func main() {
	var n int = 99999999
	fmt.Println(tab(n))
}
