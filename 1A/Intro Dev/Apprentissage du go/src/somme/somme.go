package main 

import "fmt"

func main() {
	var tint [12]int = [12]int {1,5,4,9,6,3,7,2,8,10,12,11}
	var s int
	for i := 0; i < len(tint); i++ {
		s = s + tint[i]
	}
	fmt.Println(s)
	fmt.Println(1+2+3+4+5+6+7+8+9+10+11+12)
}
		
