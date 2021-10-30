package main

import "fmt"

func main() {
	var n int = 12
	var p int = 2
	for i := 0; i < n; i++ {
		fmt.Println("L'itération", i, "commence")
		if i == p {
			continue
		}
		fmt.Println("L'itération", i, "termine")
	}
}
