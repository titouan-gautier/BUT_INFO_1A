package main

import (
	"fmt"
	"flag"
)

func tab(x int, y int, z int,tint []int) []int {
	for i := 0; i <= y; i++ {
		var prod int = x*i
		if prod%z != 0 {
			tint = append(tint,prod)
		}
	}	
	return tint
}


func main() {
	var x,y,z int
	var tint []int = make([]int, 0, y+1)
	flag.IntVar(&x,"x",0,"...")
	flag.IntVar(&y,"y",0,"...")
	flag.IntVar(&z,"z",0,"...")
	flag.Parse()
	fmt.Println(tab(x,y,z,tint))
}
