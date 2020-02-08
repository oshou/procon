package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	area1 := a * b
	area2 := c * d
	if area1 > area2 {
		fmt.Println(area1)
	} else {
		fmt.Println(area2)
	}
}
