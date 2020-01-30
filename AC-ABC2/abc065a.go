package main

import "fmt"

func main() {
	var x, a, b int
	fmt.Scan(&x, &a, &b)
	switch {
	case b-a <= 0:
		fmt.Println("delicious")
	case b-a <= x:
		fmt.Println("safe")
	default:
		fmt.Println("dangerous")
	}
}
