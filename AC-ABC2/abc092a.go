package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(int(math.Min(a, b) + math.Min(c, d)))
}
