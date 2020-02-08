package main

import (
	"fmt"
	"math"
)

func main() {
	var x, a, b float64
	fmt.Scan(&x, &a, &b)
	if math.Abs(x-a) > math.Abs(x-b) {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}
