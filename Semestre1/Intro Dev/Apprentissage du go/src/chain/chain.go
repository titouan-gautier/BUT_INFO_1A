package main

import "fmt"

func decoupe (x string) (string,string) {
	var espacePos int 
	for i := 0; i < len(x); i++ {
		if x[i] == ' ' {
			break
		} else {
			espacePos = espacePos + 1
		}
	}
	return x[:espacePos], x[espacePos+1:]
}

func main () {
	var s string = "Ordinateur Table"
	var s1, s2 string
	s1, s2 = decoupe(s)
	fmt.Print(":",s1,":\n")
	fmt.Print(":",s2,":\n")
}
