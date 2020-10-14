package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	switch {
	case a == b:
		fmt.Println(c)
	case b == c:
		fmt.Println(a)
	default:
		fmt.Println(b)
	}
}
