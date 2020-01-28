package main

import (
	"fmt"
	"math"
)

func main() {
	var f1, f2 float64
	fmt.Scan(&f1, &f2)
	fmt.Println(int(math.Abs(f1 - f2)))
}
