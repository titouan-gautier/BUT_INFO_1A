package main

import (
	"fmt"
)

func reverse1(in [3]int) (out []int) {
	out = make([]int,0)
	for cpt:=len(in)-1; cpt>=0; cpt-- {
		out = append(out, in[cpt])
	}
	return out

}

func reverse2(in [3]int) (out []int) {
	out = in[:]
	for cpt:=len(in)-1; cpt>=0; cpt-- {
		out = append(out, in[cpt])
	}
	return out

}

/* func print(in []int) {
	fmt.Printf("%c",'[')
	for cpt:=0; cpt<len(in)-1; cpt++ {
		fmt.Print(in[cpt], " ")
	}
	fmt.Print(in[len(in)-1])
	fmt.Printf("%c\n", ']') 
} */

func main() {
	tab0 := [3]int{3,25,42}
	print(tab0[:])
	tab1 := reverse1(tab0)
	print(tab1[:])
	tab2 := reverse2(tab0)
}
