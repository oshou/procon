package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var m, c, us int
	var u float64
	fmt.Scan(&n)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}

	for i := 0; i < n; i++ {
		m += abs(x[i])
		c = max(c, abs(x[i]))
		us += x[i] * x[i]
	}
	u = math.Sqrt(float64(us))

	fmt.Println(m)
	fmt.Println(u)
	fmt.Println(c)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}

}
