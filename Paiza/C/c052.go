package main

import "fmt"

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	var h, w, dx, dy int
	fmt.Scan(&h, &w)
	fmt.Scan(&dy, &dx)
	fmt.Println(abs(dx)*h + abs(dy)*w - abs(dx*dy))
}
