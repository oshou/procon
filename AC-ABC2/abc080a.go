package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	priceA := a * n
	priceB := b
	if priceA < priceB {
		fmt.Println(priceA)
	} else {
		fmt.Println(priceB)
	}
}
