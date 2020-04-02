package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, C, L, S float64
	fmt.Scanf("%f %f %f", &a, &b, &C)
	h := math.Sin(math.Pi*C/180) * b
	S = a * h / 2
	c = math.Tan(math.Pi*C/180) * a
	L = a + b + c
	fmt.Println(S)
	fmt.Println(L)
	fmt.Println(h)
}
