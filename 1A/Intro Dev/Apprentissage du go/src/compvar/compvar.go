package main

import "fmt"

func main() {
	var a int = 20
	var b int = 6
	var c int = 4
	if a > b {
		if a > c {
			fmt.Println("La variable a est la plus grande :",a)
		}
	} else if b > c {
		fmt.Println("La variable b est la plus grande :",b)
	} else {
		fmt.Println("La variable c est la plus grande :",c)
	}
		

}
