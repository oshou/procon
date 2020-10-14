package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k float64
	fmt.Scan(&n, &k)
	m := n / k
	ans1 := abs(n - math.Floor(m)*k)
	ans2 := abs(n - math.Ceil(m)*k)
	fmt.Println(ans1, ans2)
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
