package main

import "fmt"

func main() {
	var n bool
	var nb int
	var compte int
	var t []int = []int{2, 2}
	var p []int = make([]int, 0, 1)
	if t == nil {
		fmt.Println(nil)
	}
	for i := 0; i < len(t); i++ {
		n = Premier(t[i])
		if n {
			nb = 0
			for k := 0; k < len(t); k++ {
				if t[i] == t[k] {
					nb = nb + 1
				}
			}
			if nb == 1 {
				p = append(p, t[i])
			} else {
				for j := 0; j < len(p); j++ {
					if t[i] != p[j] {
						compte = compte + 1
					}
				}
				fmt.Println(compte)
				if compte == 1 {
					p = append(p, t[i])
				}
			}
		}
	}
	if p == nil {
		p = make([]int, 0)
	}
	fmt.Println(p)
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
