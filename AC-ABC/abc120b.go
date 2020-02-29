package main

import (
	"fmt"
	"os"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	var a, b, num, k int
	fmt.Scan(&a, &b, &k)
	lim := min(a, b)
	for i := lim; i >= 1; i-- {
		if a%i == 0 && b%i == 0 {
			num++
			if num == k {
				fmt.Println(i)
				os.Exit(0)
			}
		}
	}
}
