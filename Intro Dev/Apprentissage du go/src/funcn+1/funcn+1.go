package main

import "fmt"

func nbtab (x []int, y int) (bool, int) {
	if len(x) < y+1 && len(x) < y+1 {
		return false,0
	} else {
		return true,y
	}
}

func main() {
	var tab []int = []int {1,2,3,4,5}
	var n int = 10
	fmt.Println(nbtab(tab,n))
}
