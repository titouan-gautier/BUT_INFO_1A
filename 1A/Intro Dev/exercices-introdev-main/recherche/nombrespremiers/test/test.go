package main

import "fmt"

func main() {
	var n bool
	var p []int
	var t []int = nil
	for i := 0; i < len(t); i++ {
		if t != nil {
			n = Premier(t[i])
			if n {
				p = append(p, t[i])
			}
			if i == len(p) && p == nil {
				p = make([]int, 0)
			}
		} else {
			p = make([]int, 0)
		}
	}
	fmt.Println(p)
	fmt.Println(len(p))
}

func Premier(n int) (prem bool) {
	if n > 1 {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
