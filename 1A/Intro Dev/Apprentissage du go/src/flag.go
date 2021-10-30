package main

import (
	"flag"
	"fmt"
)

func main() {
	var n int
	flag.IntVar(&n, "setn", 4, "Fixe la valeur de n")
	flag.Parse()
	fmt.Println("n vaut", n)
}
