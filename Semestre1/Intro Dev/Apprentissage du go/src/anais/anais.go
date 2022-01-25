package main

import "fmt"

func main() {
	var compte float64 = 100
	var age int
	const objectif float64 = 1500
	for compte < objectif {
		age++
		var interet float64 = 3.5 * compte / 100
		var don float64 = 10 * float64(age)
		compte = compte + interet + don
		fmt.Println("Age", age ,"Compte", compte)
	}
fmt.Println(age)
}
