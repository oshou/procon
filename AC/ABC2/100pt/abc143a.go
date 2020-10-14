package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	ans := a - 2*b
	if ans < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(a - 2*b)
	}
}
