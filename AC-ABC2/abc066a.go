package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	min := math.Min(a+b, a+c)
	min = math.Min(min, b+c)
	fmt.Println(int(min))
}
