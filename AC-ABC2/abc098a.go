package main

import (
	"fmt"
)

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	add := a + b
	sub := a - b
	mul := a * b
	fmt.Println(Max(Max(add, sub), mul))
}
