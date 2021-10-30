package main

import "fmt"

func main() {
	var a int = 2
	var b int = 0
	if a > 0 && b > 0 {
		fmt.Println("Le Produit des deux variables est positif")
	} else if a < 0 || b < 0  {
		fmt.Println("Le Produit des deux variables est négatif")
	} else {
		fmt.Println("Le Produit des deux variables est nul")
	}
}	
		
