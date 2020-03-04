package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = 0
		for j := 0; j < m; j++ {
			c[i] += a[i][j] * b[j]
		}
		fmt.Println(c[i])
	}
}
