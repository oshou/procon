package main

import "fmt"

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(min(n%k, abs(n%k-k)))
}
