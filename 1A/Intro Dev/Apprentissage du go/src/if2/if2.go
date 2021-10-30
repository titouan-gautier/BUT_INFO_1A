package main

import "fmt"

func main() {
	var tint []int = []int {6,4,8,2,1}
	for i :=0; i < len(tint); i++ {
		if tint[i] < 5 {
			fmt.Println("Le nombre ",tint[i]," est plus petit que 5")
		}
	}		
}
