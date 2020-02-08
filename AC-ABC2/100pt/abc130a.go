package main

import "fmt"

func main() {
	var x, a int
	fmt.Scan(&x, &a)
	switch {
	case x < a:
		fmt.Println(0)
	default:
		fmt.Println(10)
	}
}
