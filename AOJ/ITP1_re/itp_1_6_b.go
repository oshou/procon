package main

import "fmt"

func main() {
	var n, num int
	var symbol string
	fmt.Scan(&n)
	s := make([]bool, 13)
	h := make([]bool, 13)
	c := make([]bool, 13)
	d := make([]bool, 13)
	for i := 1; i <= n; i++ {
		fmt.Scan(&symbol, &num)
		switch symbol {
		case "S":
			s[num-1] = true
		case "H":
			h[num-1] = true
		case "C":
			c[num-1] = true
		case "D":
			d[num-1] = true
		}
	}
	for i := 0; i < 13; i++ {
		if !s[i] {
			fmt.Println("S", i+1)
		}
	}
	for i := 0; i < 13; i++ {
		if !h[i] {
			fmt.Println("H", i+1)
		}
	}
	for i := 0; i < 13; i++ {
		if !c[i] {
			fmt.Println("C", i+1)
		}
	}
	for i := 0; i < 13; i++ {
		if !d[i] {
			fmt.Println("D", i+1)
		}
	}
}
