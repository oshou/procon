package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)
	if int(n)%2 == 0 {
		fmt.Println(1 / 2)
	} else {
		fmt.Println(math.Ceil(n/2) / n)
	}
}
