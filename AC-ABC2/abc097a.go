package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	direct := math.Abs(a-c) <= d
	indirect := math.Abs(a-b) <= d && math.Abs(b-c) <= d
	if direct || indirect {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
