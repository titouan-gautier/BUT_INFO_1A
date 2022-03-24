package main	

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var v int
	var e error
	v , e = strconv.Atoi(os.Args[3]) 
	if e == nil {
		fmt.Println("l'Argument", v)
	}	
	fmt.Println(e)
}

