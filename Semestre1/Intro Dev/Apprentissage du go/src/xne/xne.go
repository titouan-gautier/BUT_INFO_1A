package main

import "fmt"

func main() {
	var x float32 = 1297
	var e float32 =50.87
	fmt.Println(calc(x,e))
}

func calc(a float32, b float32) (bool,float32) {
	var i int = 0
	var res float32 = a
			for a < b {
				res = res * a
				i++
			}
	return true, res
}
