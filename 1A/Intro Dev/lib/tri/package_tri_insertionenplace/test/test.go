package main

import "fmt"

func main() {
	var i int = 1
	var t []int = []int{5, 4, 3, 2, 1}
	for i < len(t) {
		v := t[i]
		j := i - 1
		for j >= 0 && t[j] > v {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = v
		i++
	}
	fmt.Println(t)
}
