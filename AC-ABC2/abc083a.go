package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	l := a + b
	r := c + d
	switch {
	case l > r:
		fmt.Println("Left")
	case l < r:
		fmt.Println("Right")
	default:
		fmt.Println("Balanced")
	}
}
