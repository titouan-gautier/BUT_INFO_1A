package main

import "fmt"

func main () {
	var tab string = ""
	var m1 string = "bonjour"
	var m2 string = "bonsoir"
	if m1 == m2 {
		fmt.Println(false)
	}
	if m1 != tab {
		if m2 != tab {
			for i := 0 ; i<len(m1) ; i++{
				if m1[i] < m2[i]  {
					fmt.Println(true)
				}
			}
		} else {
			fmt.Println(false)
		}
	} else {
		fmt.Println(true) 
	}
}