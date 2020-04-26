package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if m[i] == 0 {
				m[i] = j
			}
		}
	}
}
