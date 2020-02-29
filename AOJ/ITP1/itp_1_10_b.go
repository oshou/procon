package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, C float64
	fmt.Scanf("%f %f %f", &a, &b, &C)
	h := math.Sin(math.Pi * C / 180)
	S := a * h / 2
	fmt.Printf("%f\n", h)
}
