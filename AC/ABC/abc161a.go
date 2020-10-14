package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	x, y = y, x
	x, z = z, x
	fmt.Println(x, y, z)
}
