package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	switch {
	case a >= 13:
		fmt.Println(b)
	case a <= 5:
		fmt.Println(0)
	default:
		fmt.Println(b / 2)
	}
}
