package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	switch x {
	case 3, 5, 7:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
