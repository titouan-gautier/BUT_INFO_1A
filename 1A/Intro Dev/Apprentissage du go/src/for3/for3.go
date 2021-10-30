package main

import "fmt"

func main() {
	var nb1 int = 14
	var nb2 int = 10
	var i int = 0
	var produit int
	for i < nb2 {
		produit = produit + nb1
		i++	
	}
	fmt.Println(produit)
}
