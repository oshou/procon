package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	x := 800 * n
	y := (n / 15) * 200
	fmt.Println(x - y)
}
