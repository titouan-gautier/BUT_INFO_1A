package main

import "fmt"

func main() {
	fmt.Println(miam(1000, 1, 1001))
}

func miam(n, m, k uint) (choco uint) {
	var e uint
	for i := n; i <= 0; i = i - m {
		n = n - m
		choco = choco + 1
		fmt.Println(n)
		e = e + 1
		if e == k {
			k = k + 1
			e = 0
		}
		if i <= 0 {
			choco = choco + e
		} else {
			miam(n, m, k)
		}
	}
	fmt.Println(n)
	fmt.Println(choco)
	return choco
}
