package main

import (
	"fmt"
)

func main() {
	var n, sum int
	div := make(map[int]bool, 0)
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		div = map[int]bool{}
		for j := 1; j*j <= i; j++ {
			if !div[j] && i%j == 0 {
				div[j] = true
				div[i/j] = true
			}
		}
		sum += i * len(div)
	}
	fmt.Println(sum)
}
