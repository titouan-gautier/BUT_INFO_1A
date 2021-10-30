package main

import "fmt"

func main() {
	var tint []int = []int {1,2,3,1,2,3,1,1,3,3,3}
	var n int = 3
	var nb int
	for i := 0; i < len(tint); i++ {
		if tint[i] == n {
			nb = nb + 1
		}
	}
	fmt.Println("L'entier n apparait ",nb," fois dans le tableau")
}

