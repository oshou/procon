package main

import (
	"fmt"
)

func gcd(a, b int) int {
	var high, low, r int
	for {
		if a > b {
			high, low = a, b
		} else {
			high, low = a, b
		}
		high = high % low
		if r == 0 {
			return low
		}
	}
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(gcd(x, y))
}
