package main

import "fmt"

func main() {
	var n float64 = 400
	if n < 300 {
		fmt.Println("Votre montant à payer est de ",n)
	} else {
		n = n * 0.95
		fmt.Println("Votre montant à payer est de ",n,"Vous avez une réduction")
	}
}
