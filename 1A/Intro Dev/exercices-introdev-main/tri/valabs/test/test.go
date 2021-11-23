package main

import "fmt"

func main() {
	var tab []int = []int{5, -22, 4, 15}
	var trier []int
	for i := 0; i < len(tab); i++ {
		trier = append(trier, tab[i])
		var j int = 0
		var v int = tab[i]
		for j < len(trier) && trier[j] <= absolue(v) {
			j++
		}
		for j < len(trier) {
			tmp := trier[j]
			trier[j] = v
			v = tmp
			j++
		}
	}
	fmt.Println(trier)
}

func absolue(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
