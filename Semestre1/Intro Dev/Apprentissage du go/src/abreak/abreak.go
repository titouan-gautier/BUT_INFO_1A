package main

import "fmt"

func main() {
	var tint []int = []int {1,2,2,4,5,3,1,1,3,3,3,7,1}
	var n int = 3
	for i:= 0; i < len(tint); i++ {
		fmt.Println(tint[i])
		if tint[i] == n {
			break
		}
	}
}
