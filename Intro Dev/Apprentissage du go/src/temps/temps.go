package main 

import "fmt"

func main () {
	var t int = 8342
	var h int = 0
	var rh int = 0
	var m int = 0
	var s int = 0
	h = t/3600
	rh = t % 3600
	m = rh/60
	s = rh%60
	fmt.Println(h,"Heure",m,"Minute",s,"Seconde")
}
