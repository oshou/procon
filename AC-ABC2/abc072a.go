package main

import (
	"fmt"
	"math"
)

func main() {
	var x, t float64
	fmt.Scan(&x, &t)
	fmt.Println(int(math.Max(x-t, 0)))
}
