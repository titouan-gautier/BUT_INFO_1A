package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var r string
	var t []string = []string{"bonjour", "salut", "bonsoir"}
	for i := 0; i < len(t); i++ {
		r = t[i]
		n = (strings.Index(r, "bon"))
	}
	fmt.Println(n)
}

/*func main() {
	str := "abcdefg"
	fmt.Println(strings.Index(str, "b"))
	//résultat:-1
}
*/
