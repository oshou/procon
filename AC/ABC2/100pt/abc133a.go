package main

import (
	"fmt"
	"math"
)

func main() {
	var n, a, b float64
	fmt.Scan(&n, &a, &b)
	fmt.Println(int(math.Min(n*a, b)))
}
