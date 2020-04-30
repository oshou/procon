package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	ad := math.Ceil(a / d)
	cb := math.Ceil(c / b)
	if ad >= cb {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
