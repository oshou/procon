package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scan(&r)
	fmt.Printf("%.6f %.6f\n", r*r*math.Pi, 2*r*math.Pi)
}
