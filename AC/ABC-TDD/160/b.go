package main

import "fmt"

func main() {
	var x, x1, x2, res int
	fmt.Scan(&x)
	x1 = x / 500
	res = x % 500
	x2 = res / 5
	fmt.Println(1000*x1 + 5*x2)
}
