package main

import "fmt"

func main() {
	var length int = 30
	var capacity int = 50
	var tab []int = make([]int, length, capacity)

	var start int = 20
	var end int = 40
	var t []int = tab[start:end]

	fmt.Println(len(tab), cap(tab))
	fmt.Println(len(t), cap(t))
}
