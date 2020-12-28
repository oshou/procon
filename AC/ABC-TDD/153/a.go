package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a int
	fmt.Scan(&h, &a)
	fmt.Println(int(math.Ceil(float64(h) / float64(a))))
}
