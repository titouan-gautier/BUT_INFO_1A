package main

import "fmt"

func maine() {
	var annee int = 2012
	if (annee%400 == 0) || (annee%4 == 0) {
		fmt.Println("Bisextile")
	} else {
		fmt.Println("Pas Bisextile")
	}

}
