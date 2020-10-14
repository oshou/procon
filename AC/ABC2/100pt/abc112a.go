package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch n {
	case 1:
		fmt.Println("Hello World")
	case 2:
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
