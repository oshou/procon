package main

import "fmt"

func main() {
	var a, b, res int
	fmt.Scan(&a, &b)
	res = a - 2*b
	if res >= 0 {
		fmt.Println(res)
	} else {
		fmt.Println(0)
	}
}
