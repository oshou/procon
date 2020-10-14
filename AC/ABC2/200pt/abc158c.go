package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	for i := 13; i <= 1000; i++ {
		if math.Floor(0.08*float64(i)) == a && math.Floor(0.1*float64(i)) == b {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
