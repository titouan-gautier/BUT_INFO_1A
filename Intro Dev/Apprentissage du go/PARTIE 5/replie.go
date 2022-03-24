package main

import "log"

func replie(t []int, f func(x int) int) {
	for i:=0 ; i<len(t) ; i++ {
		t[i] = f(t[i])
	}
}

func main() {

	var t []int = []int{1,2,3,4,5}
	var sum int

	replie(t, func (x int) int {
		sum = sum + x
		return sum
	})
	log.Print(sum)
		
}