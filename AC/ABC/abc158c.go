package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	if b < a {
		fmt.Println(-1)
		return
	}
	for i := 13; float64(i) < 12.5*(a+1); i++ {
		if math.Floor(0.08*float64(i)) == a && math.Floor(0.1*float64(i)) == b {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
