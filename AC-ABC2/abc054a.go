package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	switch {
	case a == b:
		fmt.Println("Draw")
	case a == 1:
		fmt.Println("Alice")
	case b == 1:
		fmt.Println("Bob")
	case a > b:
		fmt.Println("Alice")
	default:
		fmt.Println("Bob")
	}
}
