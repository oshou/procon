package main

import "fmt"

func main() {
	var x, n, tmp, diff int
	fmt.Scan(&x, &n)
	p := make(map[int]bool, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		p[tmp] = true
	}
	for {
		if !p[x-diff] {
			fmt.Println(x - diff)
			return
		}
		if !p[x+diff] {
			fmt.Println(x + diff)
			return
		}
		diff++
	}
}
